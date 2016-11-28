package exp

// using tree data structure to parse and evaluate mathematical expressions
// precedence by power, division, multiplication, addition, and subtraction

import (
	"fmt"
	"math"
	"strconv"

	"github.com/dockerian/go-coding/ds/str"
	u "github.com/dockerian/go-coding/utils"
)

var (
	/* see https://golang.org/ref/spec#Operators
		precedence    operator
			1             ||
			2             &&
	    3             ==  !=  <  <=  >  >=
			4             +  -  |  ^
			5             *  /  %  <<  >>  &  &^
	*/
	opPriority = map[string]int{
		"(":  0,
		"||": 11,
		"&&": 21,
		"+":  41,
		"-":  42,
		"|":  43,
		"^":  44,
		"*":  51,
		"/":  52,
		"%":  53,
		"<<": 54,
		">>": 55,
		"&":  56,
		"&^": 57,
	}
	opByLength []string
)

func init() {
	i := 0
	opByLength = make([]string, len(opPriority))
	for op := range opPriority {
		opByLength[i] = op
		i++
	}
	str.ByLength.Sort(opByLength)
}

// Exp struct
type Exp struct {
	context   string
	operators *OpStack
	operands  *OpStack
	opTree    *OpItem
	tokens    []string
}

// New returns a new instance of Exp struct
func New(s string) *Exp {
	exp := &Exp{context: s}
	exp.init()
	return exp
}

// Eval returns calculated result of the expression
func (e *Exp) Eval() float64 {
	e.init()
	result, _ := e.eval(e.opTree)
	return result
}

// String func
func (e *Exp) String() string {
	return e.toString()
}

func (e *Exp) buildParseNode() *OpItem {
	op, ob, oa := e.operators.PopItem(), e.operands.PopItem(), e.operands.PopItem()
	// u.Debug("build node: op = %+v, b = %+v, a = %+v\n", op, ob, oa)
	so, sp, sb, sa := "", "", "", ""
	if oa != nil {
		sa = oa.Expression
		oa.Parent = op
	}
	if ob != nil {
		sb = ob.Expression
		ob.Parent = op
	}
	if op != nil {
		sp = op.Expression
		so = op.Op
	}
	exprText := fmt.Sprintf("(%v %v %v)", sa, sp, sb)
	op.Left, op.Right, op.Op, op.Expression = oa, ob, so, exprText
	return op
}

func (e *Exp) buildParseTree() *OpItem {
	groupCount := 0

	for _, token := range e.tokens {
		peekToken := e.operators.Peek()
		peekPriority, _ := opPriority[peekToken]
		priority, isOperator := opPriority[token]

		switch {
		case token == "(":
			e.operators.Push(token)
			groupCount++
		case token == ")":
			if groupCount > 0 {
				peek := e.operators.Peek()
				for peek != "(" && peek != "" {
					item := e.buildParseNode()
					e.operands.PushItem(item)
					peek = e.operators.Peek()
				}
				e.operators.Pop()
				groupCount--
			}
		case isOperator:
			if e.operators.IsEmpty() || peekPriority < priority {
				e.operators.Push(token)
			} else {
				for peekToken != "" && priority < peekPriority {
					item := e.buildParseNode()
					e.operands.PushItem(item)
					peekToken = e.operators.Peek()
					peekPriority, _ = opPriority[peekToken]
					if peekToken == "(" {
						peekToken = ""
					}
				}
				e.operators.Push(token)
			}
		default:
			e.operands.Push(token)
		}
	}

	for !e.operators.IsEmpty() {
		item := e.buildParseNode()
		e.operands.PushItem(item)
	}

	return e.operands.PopItem()
}

func (e *Exp) calc(op, a, b string) (float64, error) {
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
	if _, ok := opPriority[op]; op != "" && !ok {
		return result, fmt.Errorf("Unknown operator '%v'", op)
	}

	return e.calcOp(op, fa, fb), nil
}

func (e *Exp) calcOp(op string, fa, fb float64) float64 {
	result := 0.0
	// u.Debug("calcOp: '%+v' (%+v, %+v)\n", op, fa, fb)

	switch op {
	case "":
		result = fb
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
	return result
}

func (e *Exp) checkOperatorAt(i int) (bool, string) {
	siz := len(e.context)
	for n := len(opByLength) - 1; n >= 0; n-- {
		op := opByLength[n]
		dx := i + len(op)
		if dx <= siz && op == e.context[i:dx] {
			return true, op
		}
	}
	return false, ""
}

func (e *Exp) eval(node *OpItem) (float64, error) {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			return e.calc("", "", node.Op)
		}
		op1, err := e.eval(node.Left)
		if err != nil {
			return 0.0, err
		}
		op2, err := e.eval(node.Right)
		if err != nil {
			return 0.0, err
		}
		return e.calcOp(node.Op, op1, op2), nil
	}
	return 0.0, nil
}

func (e *Exp) getOperandAt(i int) string {
	siz := len(e.context)
	var j int
	for j = i + 1; j < siz; j++ {
		ch := e.context[j]
		isOperator, _ := e.checkOperatorAt(j)
		if isOperator || ch == '(' || ch == ')' || e.isSpace(ch) {
			break
		}
	}
	if 0 <= i && j <= siz {
		return e.context[i:j]
	}
	return ""
}

func (e *Exp) init() {
	if e.operators == nil {
		u.Debug("context: '%v' ---->\n", e.context)
		e.operators = &OpStack{}
		e.operands = &OpStack{}
		e.tokens = e.tokenize()
		e.opTree = e.buildParseTree()
	}
}

func (e *Exp) isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

// tokenize converts expression context to tokens
func (e *Exp) tokenize() []string {
	siz := len(e.context)
	tokens := make([]string, 0, 1)

	for i := 0; i < siz; i++ {
		ch := e.context[i]
		switch {
		case e.isSpace(ch):
			continue
		case ch == '(':
			tokens = append(tokens, "(")
		case ch == ')':
			tokens = append(tokens, ")")
		default:
			isOperator, op := e.checkOperatorAt(i)
			if !isOperator {
				op = e.getOperandAt(i)
			}
			tokens = append(tokens, op)
			i = i + len(op) - 1
		}
	}
	return tokens
}

func (e *Exp) toString() string {
	e.init()
	if e.opTree != nil {
		return e.opTree.Expression
	}
	return ""
}
