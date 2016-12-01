// +build all pkg stock test

package stock

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	u "github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

var (
	activeRate      = 80 // percentage of active stocks
	discountRate    = 60
	exchangeNumbers = 10
	stockNumbers    = 10
	stackShift      = 2.0
	maxAmount       = 1000
	minAmount       = 100
)

// StockTestCase struct
type StockTestCase struct {
	StockID   int
	BuyAmount int
}

// TestGetBuyingStocks tests eval func
func TestGetBuyingStocks(t *testing.T) {
	market := buildStockMarket()
	var tests = []StockTestCase{
		{1, 300},
		{11 % stockNumbers, 235},
		{33 % stockNumbers, 400},
		{55 % stockNumbers, 600},
		{77 % stockNumbers, 835},
		{99 % stockNumbers, 1000},
	}

	for index, test := range tests {
		bprice := 0.0
		stocks := market.GetBuyingStocks(test.StockID, test.BuyAmount)
		u.Debug("Buying:\n%v", Buying(stocks))
		for _, stock := range stocks {
			bprice = bprice + stock.Pay
		}
		var msg = fmt.Sprintf("Expecting {%v, %4d} ==> %v", test.StockID, test.BuyAmount, bprice)
		t.Logf("Test %v: %v\n", index+1, msg)

		for i, exc := range market.Exchanges {
			mktPrice := market.GetPriceByAmount(exc.ID, test.StockID, test.BuyAmount)
			if mktPrice > 0 {
				expect := fmt.Sprintf("Expecting %s <= %v [EX: %v]", msg, mktPrice, i)
				assert.True(t, bprice <= mktPrice, expect)
			}
		}
	}
}

func buildStockMarket() *Market {
	baseStocks := make([]Stock, stockNumbers)
	exchanges := make(map[int]*Exchange, exchangeNumbers)
	market := &Market{}

	for n := 1; n < stockNumbers; n++ {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		val := rnd.Float64()*float64(500) + 1 // random range between [1.0, 501.0)
		baseStocks[n] = Stock{
			id:    n,
			stack: STACKSIZE,
			price: val,
		}
	}

	for i := 1; i < exchangeNumbers; i++ {
		stocks := make(map[int]*Stock, stockNumbers)

		for n := 1; n < stockNumbers; n++ {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			amt := rnd.Intn(maxAmount) + (maxAmount - minAmount)
			num := rnd.Intn(100)
			val := rnd.Float64()                    // random range between [0.0, 1.0)
			dta := rnd.Float64()*(stackShift-1) + 1 // random shift for stackDiff
			dis := 1 - rnd.Intn(100)/discountRate

			stocks[n] = &Stock{
				id:        n,
				active:    num/activeRate == 0,
				amount:    amt,
				stack:     STACKSIZE,
				stackDiff: baseStocks[n].stackDiff + dta*float64(dis),
				price:     baseStocks[n].price + val,
			}
		}

		exchanges[i] = &Exchange{
			ID:     i,
			Stocks: stocks,
		}
	}

	market.Exchanges = exchanges
	u.Debug("%v\n", market)

	return market
}
