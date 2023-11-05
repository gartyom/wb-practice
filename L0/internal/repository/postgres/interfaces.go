package postgres

import "github.com/gartyom/wb-practice/L0/internal/model"

type (
	PostgresOrderRepositoryInteface interface {
		GetAll() ([]model.Order, error)
		Save(id string, data []byte) error
	}
)
