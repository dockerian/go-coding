package stock

import "sort"

// ByStock is type of func as Less of sort.Interface
type ByStock func(s1, s2 *Stock) bool

// Sort func
func (by ByStock) Sort(stocks []*Stock) {
	stockSorter := &stockSorter{
		by:     by,
		stocks: stocks,
	}
	sort.Sort(stockSorter)
}

// ByStockID returns a ByStock / Less function of sort.Interface
func ByStockID() ByStock {
	return ByStock(func(s1, s2 *Stock) bool {
		return s1.id < s2.id
	})
}

// ByStockPrice returns a ByStock / Less function of sort.Interface
func ByStockPrice() ByStock {
	return ByStock(func(s1, s2 *Stock) bool {
		return s1.price < s2.price
	})
}

// ==== stockSorter type and methods ====

// stockSorter implements sort.Interface
type stockSorter struct {
	by     ByStock // closure for Less method
	stocks []*Stock
}

// Len implements Len of sort.Interface
func (ss *stockSorter) Len() int {
	return len(ss.stocks)
}

// Less implements Less of sort.Interface
func (ss *stockSorter) Less(i, j int) bool {
	return ss.by(ss.stocks[i], ss.stocks[j])
}

// Swap implments Swap of sort.Interface
func (ss *stockSorter) Swap(i, j int) {
	ss.stocks[i], ss.stocks[j] = ss.stocks[j], ss.stocks[i]
}
