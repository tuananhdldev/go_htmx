package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Setup() {

	var err error
	db, err = sql.Open("postgres", "host=0.0.0.0 port=5432 user=todo password=todo1234 dbname=todo sslmode=disable")
	if err != nil {
		fmt.Println("Could not open database", err)

	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping to database", err)
	}
}
