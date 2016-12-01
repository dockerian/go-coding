package stock

import (
	"fmt"
	"strings"

	u "github.com/dockerian/go-coding/utils"
)

// Problem:
// Deisgn a system to buy n stocks from m numbers of stock exchanges
// - each exchange may have x numbers of active stocks to sell
// - each stock prices are stacked in each stock exchange, e.g.
//   * for 1st stack (e.g. 100) the price is $50.0
//   * for 2nd stack (next 100) the price is $50.1, and etc.
//   * assume the stack amount is always 100
// The goal is to buy in lowest price

// Buyer interface
type Buyer interface {
	GetBuyingStocks(id, amount int) []BuyingStock
}

// Buying struct
type Buying []BuyingStock

// BuyingStock struct
type BuyingStock struct {
	StockID    int
	ExchangeID int
	Amount     int
	UnitPrice  float64
	StackDiff  float64
	Pay        float64
}

// String func for Buying
func (bu Buying) String() string {
	header := (&BuyingStock{}).StringHeader()
	sTable := ""
	for _, stock := range []BuyingStock(bu) {
		sTable = fmt.Sprintf("%s%v\n", sTable, &stock)
	}
	return fmt.Sprintf("%s\n%s\n", header, sTable)
}

// String func for BuyingStock
func (bs *BuyingStock) String() string {
	price := u.FmtComma(fmt.Sprintf("%.2f", bs.Pay))
	return fmt.Sprintf("%8d\t%5d\t%8d\t%12s\t%6.2f\t%6.2f",
		bs.ExchangeID, bs.StockID, bs.Amount, price, bs.UnitPrice, bs.StackDiff)
}

// StringHeader func for BuyingStock
func (bs *BuyingStock) StringHeader() string {
	splits := strings.Repeat("-", 70)
	header := fmt.Sprintf("%8s\t%5s\t%8s\t%12s\t%6s\t%6s\n",
		"EX", "ID", "Amount", "Pay", "U.Price", "Stack")
	return header + splits
}
