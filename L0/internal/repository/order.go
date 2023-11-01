package repository

import (
	"github.com/gartyom/wb-practice/L0/internal/cache"
)

type orderRepository struct {
	cache *cache.Cache
}

func NewOrderRepository(cch *cache.Cache) OrderRepositoryInteface {
	return &orderRepository{
		cache: cch,
	}
}

func (r *orderRepository) GetById(id string) ([]byte, error) {
	data, err := r.cache.GetById(id)
	return data, err
}
