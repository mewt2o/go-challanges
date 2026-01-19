package main

import (
	"math/rand"
	"time"
)

type StockPrice struct {
	Symbol string
	Price  float64
	Source string
}

type StockSource struct {
	Name  string
	Delay time.Duration
}

func (s StockSource) FetchPrice(symbol string) StockPrice {
	time.Sleep(s.Delay)

	basePrice := 100.0
	variation := rand.Float64()*10 - 5

	return StockPrice{
		Symbol: symbol,
		Price:  basePrice + variation,
		Source: s.Name,
	}
}

func GetSources() []StockSource {
	return []StockSource{
		{Name: "AlphaAPI", Delay: 200 * time.Millisecond},
		{Name: "BetaFinance", Delay: 150 * time.Millisecond},
		{Name: "GammaStocks", Delay: 300 * time.Millisecond},
		{Name: "DeltaMarket", Delay: 100 * time.Millisecond},
	}
}
