package ctx_in_http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Start() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// this request has a context
	// the context ensures that if the request takes longer than 2s, it will be cancelled.
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/todos", nil)
	if err != nil {
		fmt.Println("Error creating request")
		return
	}

	resp := callAPI(req)
	_ = resp
	// process response
}

func callAPI(req *http.Request) []byte {
	select {
	case <-req.Context().Done():
		fmt.Println("Request context expired: ", req.Context().Err())
		return nil
	}
}
