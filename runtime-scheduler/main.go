package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
}

func main() {
	// Start an HTTP server for pprof
	go func() {
		fmt.Println("Starting pprof server at http://localhost:6060")
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup

	// Simulate concurrent workload
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go work(&wg, i)
	}

	wg.Wait()

	// Prevent the program from exiting immediately after the work is done
	// Keep the server running for profiling
	fmt.Println("All goroutines finished. Press Ctrl+C to stop the profiler.")
	select {} // Block the program from exiting by using an empty select (infinite wait)
}
