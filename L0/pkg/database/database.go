package database

import (
	"fmt"
	"log"

	"database/sql"
	"github.com/gartyom/wb-practice/L0/internal/config"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	pg, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = pg.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database successfully connected")

	return pg
}
