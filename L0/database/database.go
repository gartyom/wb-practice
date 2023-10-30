package database

import (
	"context"
	"fmt"

	"github.com/gartyom/wb-pratctice/L0/config"
	"github.com/jackc/pgx/v5"
)

func Connect(cfg *config.Config) *pgx.Conn {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	pg, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		panic(err)
	}

	err = pg.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	return pg
}
