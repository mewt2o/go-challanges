package main

import (
	"context"
	"fmt"
	"time"
)

func AggregateStockPrice(symbol string, sources []StockSource, timeout time.Duration) (StockPrice, error) {
	// TODO: Implement this function
	//
	// Hints:
	// 1. Create a channel to receive stock prices
	// 2. Launch a goroutine for each source
	// 3. Use select with time.After for timeout
	// 4. Return the first result received
	// 5. Consider using context for cancellation

	panic("not implemented")
}

func AggregateStockPriceWithContext(ctx context.Context, symbol string, sources []StockSource) (StockPrice, error) {
	// TODO: Implement this function
	//
	// Hints:
	// 1. Similar to above but use ctx.Done() instead of time.After
	// 2. This allows external cancellation control

	panic("not implemented")
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
