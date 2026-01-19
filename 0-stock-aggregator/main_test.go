package main

import (
	"context"
	"testing"
	"time"
)

func TestAggregateStockPrice_Success(t *testing.T) {
	sources := GetSources()

	price, err := AggregateStockPrice("AAPL", sources, 2*time.Second)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if price.Symbol != "AAPL" {
		t.Errorf("Expected symbol AAPL, got: %s", price.Symbol)
	}

	if price.Source == "" {
		t.Error("Expected source name, got empty string")
	}

	if price.Source != "DeltaMarket" {
		t.Logf("Warning: Expected fastest source DeltaMarket, got: %s", price.Source)
	}
}

func TestAggregateStockPrice_Timeout(t *testing.T) {
	slowSources := []StockSource{
		{Name: "SlowAPI", Delay: 2 * time.Second},
	}

	_, err := AggregateStockPrice("AAPL", slowSources, 100*time.Millisecond)

	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
}

func TestAggregateStockPrice_Performance(t *testing.T) {
	sources := GetSources()

	start := time.Now()
	price, err := AggregateStockPrice("AAPL", sources, 2*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if elapsed > 200*time.Millisecond {
		t.Errorf("Expected response in < 200ms, took: %v", elapsed)
	}

	t.Logf("Got price $%.2f from %s in %v", price.Price, price.Source, elapsed)
}

func TestAggregateStockPriceWithContext_Cancellation(t *testing.T) {
	sources := GetSources()
	ctx, cancel := context.WithCancel(context.Background())

	cancel()

	_, err := AggregateStockPriceWithContext(ctx, "AAPL", sources)

	if err == nil {
		t.Fatal("Expected context cancellation error, got nil")
	}
}

func TestAggregateStockPriceWithContext_Success(t *testing.T) {
	sources := GetSources()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	price, err := AggregateStockPriceWithContext(ctx, "AAPL", sources)

	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if price.Symbol != "AAPL" {
		t.Errorf("Expected symbol AAPL, got: %s", price.Symbol)
	}
}

func BenchmarkAggregateStockPrice(b *testing.B) {
	sources := GetSources()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := AggregateStockPrice("AAPL", sources, 2*time.Second)
		if err != nil {
			b.Fatal(err)
		}
	}
}
