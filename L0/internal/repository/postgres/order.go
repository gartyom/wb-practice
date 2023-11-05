package postgres

import (
	"database/sql"

	"github.com/gartyom/wb-practice/L0/internal/model"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) PostgresOrderRepositoryInteface {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Save(id string, orderData []byte) error {
	_, err := r.db.Exec("INSERT INTO \"order\" VALUES($1, $2)", id, orderData)
	return err
}

func (r *orderRepository) GetAll() ([]model.Order, error) {
	rows, err := r.db.Query("SELECT * FROM \"order\"")
	if err != nil {
		return nil, err
	}

	result := make([]model.Order, 0)

	for rows.Next() {
		var order model.Order

		if err := rows.Scan(&order.OrderUID, &order.Data); err != nil {
			return nil, err
		}

		result = append(result, order)
	}

	return result, nil
}
