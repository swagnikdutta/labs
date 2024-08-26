package ctx_with_deadline

import (
	"context"
	"fmt"
	"time"
)

func Start() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	// Do we need this cancel? What's the deadline for then?
	//
	// Not needed. Could have written,
	// ```go
	//     ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	//     go performTask(ctx)
	// ```
	// The returned context's Done channel is closed when the deadline expires,
	// when the returned cancel function is called, or when the parent context's Done channel is closed,
	// whichever happens first.

	go performTask(ctx)

	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Deadline must have exceeded: ", ctx.Err())

		return // seems optional
	}
}
