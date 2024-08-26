package cancelling_ctx

import (
	"context"
	"fmt"
	"time"
)

func Start() {
	ctx, cancel := context.WithCancel(context.Background())

	go performTask(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	// This sleep is needed to print that "Task cancelled" on line 8
	time.Sleep(1 * time.Second)
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			fmt.Println("Performing task...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
