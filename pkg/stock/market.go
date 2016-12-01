package stock

import (
	"fmt"
	"math"
	"sort"

	"github.com/dockerian/go-coding/ds/mathEx"
	u "github.com/dockerian/go-coding/utils"
)

// Market struct
type Market struct {
	Exchanges map[int]*Exchange
}

// ==== Market methods ====

// GetBuyingStocks returns a set of stock to buy at the lowest price
func (m *Market) GetBuyingStocks(id, amount int) []BuyingStock {
	choices := make(map[int]float64) // exchangeID => unit price
	stacked := make(map[int]int)     // exchangeID => stack
	amounts := make(map[int]int)     // exchangeID => stock amount
	selects := make(map[int]int)     // exchangeID => buying amount
	results := make([]BuyingStock, 0, 1)

	var amountLeft = amount

	for exID, ex := range m.Exchanges {
		if stock, found := ex.Stocks[id]; found {
			if stock.active && stock.amount > 0 {
				amounts[exID] = stock.amount
				choices[exID] = stock.price
				stacked[exID] = 1
			}
		}
	}

	for amountLeft > 0 {
		var exchangeID int
		outStock, minPrice := true, math.MaxFloat64
		for exID, price := range choices {
			amountInStock := amounts[exID]
			if price < minPrice && amountInStock > 0 {
				minPrice = price
				exchangeID = exID
				outStock = false
			}
		}
		if outStock {
			u.Debug("Check stock ID = %v [out of stock]\n", id)
			break
		}

		price := choices[exchangeID]
		stack := stacked[exchangeID] + 1
		amountInStock := amounts[exchangeID]
		x := mathEx.MinInt(STACKSIZE, amountLeft, amountInStock)
		amounts[exchangeID] = amountInStock - x
		selectAmount, _ := selects[exchangeID]
		selects[exchangeID] = selectAmount + x
		amountLeft = amountLeft - x
		u.Debug("Found stock ID = %v : %3d x %6.2f {EX: %3d} +%v stack\n", id, x, price, exchangeID, stack-1)
		choices[exchangeID] = m.GetUnitPriceByStack(exchangeID, id, stack)
		stacked[exchangeID] = stack
	}

	for exchangeID, amount := range selects {
		stock := m.Exchanges[exchangeID].Stocks[id]
		buyingStock := BuyingStock{
			StockID:    id,
			ExchangeID: exchangeID,
			Amount:     amount,
			Pay:        m.GetPriceByAmount(exchangeID, id, amount),
			StackDiff:  stock.stackDiff,
			UnitPrice:  stock.price,
		}
		results = append(results, buyingStock)
	}

	return results
}

// GetExchanges returns sorted exchanges list
func (m *Market) GetExchanges() []*Exchange {
	i, exchanges := 0, make([]*Exchange, len(m.Exchanges))
	for _, ex := range m.Exchanges {
		exchanges[i] = ex
		i++
	}
	sort.Sort(byExchangeID(exchanges))
	return exchanges
}

// GetPriceByAmount returns the total per exchange, stock, and number of stocks
func (m *Market) GetPriceByAmount(exchangeID, stockID, amount int) float64 {
	if exchange, okay := m.Exchanges[exchangeID]; okay {
		if stock, found := exchange.Stocks[stockID]; found {
			if stock.active {
				return stock.GetPriceByAmount(amount)
			}
		}
	}
	return 0.0
}

// GetUnitPriceByStack returns the price for spcified exchange, stock, and stack
func (m *Market) GetUnitPriceByStack(exchangeID, stockID, stack int) float64 {
	if exchange, okay := m.Exchanges[exchangeID]; okay {
		if stock, found := exchange.Stocks[stockID]; found {
			if stock.active {
				return stock.GetUnitPriceByStack(stack)
			}
		}
	}
	return 0.0
}

// String func for Market
func (m *Market) String() string {
	eTable := ""
	exchanges := m.GetExchanges()
	for _, exchange := range exchanges {
		eTable = fmt.Sprintf("%s%v\n", eTable, exchange)
	}
	return eTable
}
