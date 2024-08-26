package ctx_with_value

import (
	"context"
	"fmt"
)

func Start() {
	// WithValue returns only the context
	ctx := context.WithValue(context.Background(), "userID", 123)
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("userID").(int)
	fmt.Println("Processing request for user: ", userID)
}
