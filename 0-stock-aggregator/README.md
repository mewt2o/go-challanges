# Challenge 0: Stock Aggregator

## Overview
**Topics**: Goroutines, Channels, Concurrency Basics  
**Estimated Time**: 30-45 minutes

## Objective
Build a stock price aggregator that fetches prices from multiple sources concurrently and returns the fastest response. This challenge teaches you the fundamentals of goroutines and channels.

## Problem Description
You need to fetch stock prices from multiple APIs concurrently. The function should:
1. Query multiple stock price sources simultaneously
2. Return the **first** response received (fastest source)
3. Handle timeouts if all sources are too slow
4. Cancel remaining requests once first response arrives

## Structure
```
0-stock-aggregator/
├── main.go           # Your solution goes here
├── main_test.go      # Test cases
├── sources.go        # Mock API sources (don't modify)
└── README.md         # This file
```

## Running the Challenge

```bash
# Run the program
go run .

# Run tests
go test -v

# Run benchmarks
go test -bench=. -benchmem

# Check test coverage
go test -cover
```