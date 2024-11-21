package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a large slice
	data := make([]int, 1_000_000)
	for i := range data {
		data[i] = i
	}

	// Measure performance of bounds checking
	start := time.Now()
	for i := 0; i < 10_000; i++ {
		for j := 0; j < len(data); j++ { // This is where bounds checking happens
			_ = data[j]
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken with bounds checking: %s\n", elapsed)
}
