package main

import (
	"fmt"
	"net/http"

	coingecko "github.com/superoo7/go-gecko/v3"
)

const bitcoinSupply = 21_000_000.0
const reportTemplate = `----------------------------
Coin Statistics for %v (%v):
Current Price: %.8f
All Time High (could be inaccurate): %.8f
Maximum satoshi value by circulating supply: %.8f
Maximum satoshi value by total supply: %.8f
----------------------------`

type coinStatistics struct {
	Name              string
	Symbol            string
	CirculatingSupply float64
	TotalSupply       float64
	CurrentPrice      float64
	AllTimeHighPrice  float64
}

func (c coinStatistics) MaxBitcoinPriceByCirculatingSupply() float64 {
	return bitcoinSupply / c.CirculatingSupply
}

func (c coinStatistics) MaxBitcoinPriceByTotalSupply() float64 {
	if c.TotalSupply == 0 {
		return 0
	}

	return bitcoinSupply / c.TotalSupply
}

func main() {
	hc := http.Client{}

	cg := coingecko.NewClient(&hc)

	cm, err := cg.CoinsMarket("BTC", nil, "market_cap_desc", 10, 1, false, nil)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, c := range *cm {
		cs := coinStatistics{
			CirculatingSupply: c.CirculatingSupply,
			CurrentPrice:      c.CurrentPrice,
			TotalSupply:       c.TotalSupply,
			Name:              c.Name,
			Symbol:            c.Symbol,
			AllTimeHighPrice:  c.ATH,
		}

		fmt.Printf(reportTemplate, cs.Name, cs.Symbol, cs.CurrentPrice, cs.AllTimeHighPrice, cs.MaxBitcoinPriceByCirculatingSupply(), cs.MaxBitcoinPriceByTotalSupply())
	}

}
