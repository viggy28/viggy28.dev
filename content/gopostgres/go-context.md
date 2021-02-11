---
date: 2020-10-02T11:00:08-04:00
description: "Understanding context in Go using Postgres"
featured_image: "/images/postgres-logo.png"
tags: ["go","postgres"]
title: "go context"
---

# Understanding context in Golang through Postgres

I was trying to learn Go package `context` especially with respect to Postgres. On a very high level `context` provides context to the operation. Yeah, I agree, the previous statement doesn't really add much value, but hold on I don't really know how to explain it, rather let's go over some code. Sometimes its easier to understand by seeing it in action :) 

The complete program is [here](https://gitlab.com/viggy28-websites/viggy28.dev/blob/master/static/files/gopostgres/go-context/main.go)

- **Time based context:**

  There are two methods. `context.WithTimeout()` `context.WithDeadline()`

  ```go
  package main
  
  import (
  	"context"
  	"database/sql"
  	"log"
  	"os"
  
  	_ "github.com/lib/pq"
  )
  
  func main() {
  	dsn := "user=postgres dbname=postgres sslmode=disable port=5433 password=replaceit"
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
  	ctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)
  	defer cancel()
  	_, err = db.QueryContext(ctx, "select * from generate_series(1, 5000000000)")
  	if err != nil {
  		log.Printf("error querying %v", err)
  	}
  	log.Println("done")
  ```

  `	ctx := context.Background()` creates an empty context. 

  Then `	ctx, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)` which creates a timed context for `ctx`. When we use that context in `db.QueryContext`  the context gets cancelled at 2s. Meaning any query taking more than 2s will be stopped.

- **Using WithCancel()**:

  ```go
  package main
  
  import (
  	"context"
  	"database/sql"
  	"log"
  	"os"
  
  	_ "github.com/lib/pq"
  )
  
  func main() {
  	dsn := "user=postgres dbname=postgres sslmode=disable port=5433 password=replaceit"
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
  	ctx, cancel := context.WithCancel(ctx)
  	go func() {
  	   	select {
  	   	case <-time.After(2000 * time.Millisecond):
  	   		fmt.Fprint(os.Stderr, "query timed out \n")
  	   		cancel()
  	   	}
  	   }()
  	_, err = db.QueryContext(ctx, "select * from generate_series(1, 50000000)")
  	cancel()
  ```

  Achieving the same thing (time out queries which takes more than 2 seconds) but using `cancel()`. In a separate go routine we checking whether its been 2 seconds and by calling `cancel()` it cancels the context.

  

  I don't know is there any preferable method. Certainly for time based contexts,  `context.WithTimeout()` or `context.WithDeadline()` seems more straightforward.