package cacher

import (
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

func (cch *Cache) GetById(id string) ([]byte, error) {
	data := cch.OrderData[id]
	if data == nil {
		return nil, errors.New("Cache: Not found")
	}
	return data, nil
}

func (cch *Cache) Save(id string, data []byte) {
	cch.OrderData[id] = data
}
