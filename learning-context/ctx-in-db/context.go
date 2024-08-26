package ctx_in_db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Start will not run, it will error. Code only for reading purpose.
func Start() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", "connection-string")
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}
	defer db.Close()

	// QueryContext executes a query that returns rows, typically a SELECT.
	// When executing a database query with db.QueryContext(), the context ensures
	// that the operation will be cancelled, if it exceeds the specified timeout.
	rows, err := db.QueryContext(ctx, "SELECT * from users")
	if err != nil {
		fmt.Println("Error executing query")
		return
	}
	defer rows.Close()

	// process query results
}
