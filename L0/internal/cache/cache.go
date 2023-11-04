package cache

import (
	"database/sql"
	"errors"
)

type Cache struct {
	OrderData map[string][]byte
}

func New() *Cache {
	return &Cache{
		OrderData: make(map[string][]byte),
	}
}

func (cch *Cache) Recover(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM \"order\"")
	if err != nil {
		return err
	}
	defer rows.Close()
	var uid string
	var data []byte
	for rows.Next() {
		if err := rows.Scan(&uid, &data); err != nil {
			return err
		}

		cch.OrderData[uid] = data
	}
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (cch *Cache) GetById(id string) ([]byte, error) {
	data := cch.OrderData[id]
	if data == nil {
		return nil, errors.New("Cache: Not found")
	}
	return data, nil
}
