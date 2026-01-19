package main

import (
	"context"
	"fmt"
	"time"
)

func AggregateStockPrice(symbol string, sources []StockSource, timeout time.Duration) (StockPrice, error) {
	ch := make(chan StockPrice, len(sources))
	for _, source := range sources {
		go func(s StockSource) {
			price := s.FetchPrice(symbol)
			ch <- price
		}(source)
	}
	select {
	case price := <-ch:
		return price, nil
	case <-time.After(timeout):
		return StockPrice{}, fmt.Errorf("timeout")
	}
}

func AggregateStockPriceWithContext(ctx context.Context, symbol string, sources []StockSource) (StockPrice, error) {
	ch := make(chan StockPrice, len(sources))
	for _, source := range sources {
		go func(s StockSource) {
			price := s.FetchPrice(symbol)
			select {
			case ch <- price:
			case <-ctx.Done():
			}
		}(source)
	}
	select {
	case price := <-ch:
		return price, nil
	case <-ctx.Done():
		return StockPrice{}, ctx.Err()
	}
}

func main() {
	sources := GetSources()

	fmt.Println("Fetching AAPL stock price from multiple sources...")

	start := time.Now()
	price, err := AggregateStockPrice("AAPL", sources, 5*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Got price $%.2f from %s in %v\n", price.Price, price.Source, elapsed)
}
