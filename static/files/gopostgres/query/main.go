package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "user=postgres dbname=postgres sslmode=disable password=mysecretpassword host=0.0.0.0"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	var (
		id           int
		team_key     string
		publisher_id int
		home_site_id sql.NullInt32
	)
	rows, err := db.Query("select * from teams where home_site_id IS NOT NULL")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &team_key, &publisher_id, &home_site_id)
		if err != nil {
			panic(err)
		}
	}
	log.Println(id, team_key, publisher_id, home_site_id)
	res, err := db.Exec("update teams set home_site_id=1 where id=1")
	if err != nil {
		panic(err)
	}
	n, err := res.RowsAffected()
	m, err := res.LastInsertId()
	log.Println(n, m)
}
