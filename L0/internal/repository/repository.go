package repository

import (
	"database/sql"

	"github.com/gartyom/wb-practice/L0/internal/repository/postgres"
)

type Repository struct {
	Order postgres.PostgresOrderRepositoryInteface
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Order: postgres.NewOrderRepository(db),
	}
}
