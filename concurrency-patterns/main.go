package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func task(b []byte) int {
	time.Sleep(10 * time.Microsecond)
	return 1
}

func tester(workers int) int64 {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	var count int64
	ch := make(chan []byte, workers)

	// start a pool of workers
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	channelOpen := true
	for channelOpen {
		select {
		case <-ctx.Done():
			close(ch)
			channelOpen = false
			break
		default:
			b := make([]byte, 1024)
			ch <- b
		}
	}
	wg.Wait()
	return count
}

func main() {
	workers := []int{2, 5, 10, 50, 100, 1000}

	for _, w := range workers {
		fmt.Printf("Count: %d, with %d workers\n", tester(w), w)
	}

}
