package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "user=postgres dbname=postgres sslmode=disable port=5433 password=Test_1234"
	log.Println(dsn)
	os.Environ()
	db, err := sql.Open("postgres", dsn)
	defer db.Close()
	if err != nil {
		log.Fatalf("error opening a connection to database %v", err)
	}
	err = db.PingContext(context.Background())
	if err != nil {
		log.Fatalf("error connecting to the database %v", err)
	}
	ctx := context.Background()
	// for timed contexts
	ctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)
	defer cancel()
	_, err = db.QueryContext(ctx, "select * from generate_series(1, 500000000)")
	select {
	case <-ctx.Done():
		fmt.Fprint(os.Stderr, "query timed out \n")
	default:
		fmt.Fprint(os.Stderr, "query finished \n")
	}
	ctx, cancel = context.WithCancel(ctx)
	go func() {
		select {
		case <-time.After(2000 * time.Millisecond):
			fmt.Fprint(os.Stderr, "query timed out \n")
			cancel()
		}
	}()
	_, err = db.QueryContext(ctx, "select * from generate_series(1, 50000000)")
	cancel()

	_, err = db.QueryContext(context.Background(), "select * from generate_series(1, 5000000000)")
	if err != nil {
		log.Printf("error querying %v", err)
	}
	log.Println("done")
}
