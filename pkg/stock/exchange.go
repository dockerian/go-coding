package stock

import "fmt"

// Exchange struct
type Exchange struct {
	ID     int
	Stocks map[int]*Stock
}

// ==== Exchange methods ====

// StockHeader func for stocks list header
func (e *Exchange) StockHeader() string {
	s := &Stock{}
	return s.StringHeader()
}

// GetStocks returns sorted stocks list
func (e *Exchange) GetStocks() []*Stock {
	i, stocks := 0, make([]*Stock, len(e.Stocks))
	for _, stock := range e.Stocks {
		stocks[i] = stock
		i++
	}
	ByStockID().Sort(stocks)
	return stocks
}

// String func for Exchange
func (e *Exchange) String() string {
	header := fmt.Sprintf("EX: %d\n========\n%s", e.ID, e.StockHeader())
	sTable := ""
	stocks := e.GetStocks()
	for _, stock := range stocks {
		sTable = fmt.Sprintf("%s%v\n", sTable, stock)
	}
	return fmt.Sprintf("%s\n%s\n", header, sTable)
}
