package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Simulate high-memory usage by allocating large objects
	for i := 0; i < 10; i++ {
		makeLargeAllocations()
		runtime.GC() // Force garbage collection
		printMemoryStats()
		time.Sleep(1 * time.Second) // Simulate workload
	}
}

func makeLargeAllocations() {
	// Allocate a large slice to simulate large heap usage
	largeSlice := make([]byte, 2*1024*1024*1024) // 2GB
	_ = largeSlice
}

func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tHeapSys = %v MiB", m.HeapSys/1024/1024)
	fmt.Printf("\tHeapReleased = %v MiB", m.HeapReleased/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
