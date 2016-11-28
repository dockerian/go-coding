package puzzle

// parse and evaluate mathematical expressions
// precedence by power, division, multiplication, addition, and subtraction

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dockerian/go-coding/ds/stack"
	u "github.com/dockerian/go-coding/utils"
)

var (
	operators = []string{
		"+",
		"-",
		"*",
		"/",
		"^",
	}
	operatorPriority = map[string]int{
		"+": 1,
		"-": 2,
		"*": 3,
		"/": 4,
		"^": 5,
	}
)

func eval(s string) (float64, error) {
	e := NewEval(s)
	return e.eval()
}

// Eval struct is a string stack implementation
type Eval struct {
	expression string
	operators  stack.Str
	operands   stack.Str
}

// NewEval return a new instance of Eval struct
func NewEval(e string) *Eval {
	return &Eval{
		expression: e,
		operators:  stack.Str{},
		operands:   stack.Str{},
	}
}

func (e *Eval) calc() (string, error) {
	op, bv, av := e.operators.Pop(), e.operands.Pop(), e.operands.Pop()
	if op == "" && av == "" {
		return bv, nil
	}
	result, err := evaluate(op, av, bv)
	u.Debug("calc: '%s' %s '%s' == %v \n", av, op, bv, result)
	if err != nil {
		return "", err
	}
	item := fmt.Sprintf("%f", result)
	return item, nil
}

func (e *Eval) eval() (float64, error) {
	tokens := strings.Split(e.expression, " ")
	u.Debug("eval: '%v'\n", e.expression)

	for _, token := range tokens {
		priority, isOperator := operatorPriority[token]
		peekToken := e.operators.Peek()
		peekPriority, _ := operatorPriority[peekToken]

		if isOperator {
			if e.operators.IsEmpty() || peekPriority < priority {
				e.operators.Push(token)
			} else {
				for !e.operators.IsEmpty() && priority < peekPriority {
					result, err := e.calc()
					if err != nil {
						return 0.0, err
					}
					e.operands.Push(result)
				}
				e.operators.Push(token)
			}
		} else {
			e.operands.Push(token)
		}
	}

	for !e.operators.IsEmpty() {
		result, err := e.calc()
		if err != nil {
			return 0.0, err
		}
		e.operands.Push(result)
	}

	peekValue := e.operands.Peek()
	result, err := strconv.ParseFloat(peekValue, 64)
	return result, err
}

func evaluate(op, a, b string) (float64, error) {
	var result float64
	if a == "" {
		a = "0.0" // allow 1st operand to be empty
	}
	fa, err1 := strconv.ParseFloat(a, 64)
	if err1 != nil {
		return result, fmt.Errorf("Cannot parse '%v' to float64: %v", a, err1)
	}
	fb, err2 := strconv.ParseFloat(b, 64)
	if err2 != nil {
		return result, fmt.Errorf("Cannot parse '%v' to float64: %v", b, err2)
	}
	if _, ok := operatorPriority[op]; !ok {
		return result, fmt.Errorf("Unknown operator '%v'", op)
	}

	switch op {
	case "+":
		result = fa + fb
	case "-":
		result = fa - fb
	case "*":
		result = fa * fb
	case "/":
		result = fa / fb
	case "^":
		result = math.Pow(fa, fb)
	default:
	}
	return result, nil
}
