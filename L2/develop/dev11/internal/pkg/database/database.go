package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func New() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=postgres password=pass dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database succesfully connected!")

	return db

}
