package main

import (
	"fmt"
	"runtime"
	"time"
)

func allocateMemory(size int) {
	// Allocate a significant amount of memory
	for i := 0; i < size; i++ {
		_ = make([]byte, 1024*1024) // Allocate 1 MB
	}
}

func main() {
	var memStats runtime.MemStats

	// Run multiple allocation rounds
	for i := 0; i < 5; i++ {
		allocateMemory(10) // Allocate 10 MB in this round
		runtime.GC()       // Manually trigger garbage collection
		runtime.ReadMemStats(&memStats)

		// Print GC stats
		fmt.Printf(
			"Round %d: Alloc = %v KiB, HeapAlloc = %v KiB, TotalAlloc = %v MiB, NumGC = %v\n",
			i+1,
			memStats.Alloc/1024,
			memStats.HeapAlloc/1024,
			memStats.TotalAlloc/1024/1024,
			memStats.NumGC,
		)

		// Sleep to allow for observable performance
		time.Sleep(2 * time.Second)
	}
}
