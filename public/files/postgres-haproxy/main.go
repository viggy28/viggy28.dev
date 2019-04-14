// function to evaluate the performance of pgbounce
// do docker compose up on docker-compose-with-pgbouncer.yml file
// assuming a postgres with haproxy setup runs on port 15432 and pgbouncer on 5432

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "Test_1234"
	dbName     = "bdrdemo"
	dbHost     = "localhost"
)

func main() {
	var method string
	for i := 0; i < 14; i++ {
		start := time.Now()
		r := new(big.Int)
		fmt.Println(r.Binomial(1000, 10))
		var dsn string
		dsn = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port="+haProxy(), dbUser, dbPassword, dbName, dbHost)
		method = "haproxy"
		if i%2 == 0 {
			dsn = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port="+pgBounce(), dbUser, dbPassword, dbName, dbHost)
			method = "pgbounce"
		}
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("FATAL connecting to database %v", err)
		}
		defer db.Close()
		err = db.Ping()
		if err != nil {
			log.Println("ERROR pinging database", err)
		}
		elapsed := time.Since(start)
		log.Printf("It took %s %s", elapsed, method)
	}
}

func pgBounce() string {
	return "5432"
}

func haProxy() string {
	return "15432"
}
