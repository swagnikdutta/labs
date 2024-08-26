package main

import (
	"fmt"

	cancellingctx "learning-context/cancelling-ctx"
	ctxindb "learning-context/ctx-in-db"
	ctxinhttp "learning-context/ctx-in-http"
	ctxwithdeadline "learning-context/ctx-with-deadline"
	ctxwithvalue "learning-context/ctx-with-value"
)

func main() {
	fmt.Println("Demo: context with value")
	ctxwithvalue.Start()

	fmt.Println("\nDemo: cancelling context")
	cancellingctx.Start()

	fmt.Println("\nDemo: context with deadline")
	ctxwithdeadline.Start()

	fmt.Println("\nDemo: context in http requests")
	ctxinhttp.Start()

	fmt.Println("\nDemo: context with database operations")
	ctxindb.Start()
}
