package stock

import (
	"fmt"
	"strings"

	"github.com/dockerian/go-coding/ds/mathEx"
)

const (
	// MAXSTACK is the maximum stacks allowed for a stack
	MAXSTACK = 10

	// STACKSIZE is the number of stocks in one stack
	STACKSIZE = 100

	// STOCKAMOUNT is the default number of stocks
	STOCKAMOUNT = 1000
)

// Stock struct represents exchangeable stock
// - assuming same price for buying and selling
// - assuming stack is always the same (for simplicity)
// - note: for varied stack, see GetUnitPriceByBoughtAmount
type Stock struct {
	id        int
	stack     int     // always = STACKSIZE
	stackDiff float64 // increasing price for each more stack
	price     float64
	active    bool
	amount    int
}

// NewStock builds a stock
func NewStock(id, amount int, price, stackDiff float64) *Stock {
	stock := &Stock{
		id:        id,
		price:     price,
		stack:     STACKSIZE,
		stackDiff: stackDiff,
		amount:    amount,
		active:    true,
	}
	return stock
}

// ==== Stock methods ====

// GetPriceByAmount returns the total price by number of stocks
func (s *Stock) GetPriceByAmount(amount int) float64 {
	stackSize := s.stack
	if s.stack <= 0 {
		stackSize = STACKSIZE
	}
	var stack = 1
	var price float64
	amountLeft := amount

	for amountLeft > 0 {
		x := mathEx.MinInt(amountLeft, stackSize)
		stackPrice := s.GetUnitPriceByStack(stack)
		price = price + stackPrice*float64(x)
		amountLeft = amountLeft - x
		stack++
	}
	return price
}

// GetUnitPriceByBoughtAmount returns the unit price per bought amount
func (s *Stock) GetUnitPriceByBoughtAmount(boughtAmount int) float64 {
	price := s.price
	stack := s.stack
	if stack <= 0 {
		stack = STACKSIZE
	}
	n := 1
	for boughtAmount > stack {
		if n > MAXSTACK {
			break
		}
		boughtAmount = boughtAmount - stack
		price = price + s.stackDiff
		n++
	}
	return price
}

// GetUnitPriceByStack returns the unit price for specified stack
func (s *Stock) GetUnitPriceByStack(stack int) float64 {
	if stack <= 0 {
		return 0.0
	}
	price := s.price
	for i := 2; i <= stack && i <= MAXSTACK; i++ {
		price = price + s.stackDiff
	}
	return price
}

// String func for Stock
func (s *Stock) String() string {
	return fmt.Sprintf("%8d\t%6.2f\t%6.2f\t%8d\t%v",
		s.id, s.price, s.stackDiff, s.amount, s.active)
}

// StringHeader func for stocks list header
func (s *Stock) StringHeader() string {
	splits := strings.Repeat("-", 60)
	header := fmt.Sprintf("%8s\t%6s\t%6s\t%8s\t%v\n",
		"ID", "Price", "Stack", "Amount", "Active")
	return header + splits
}
