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

func (r *orderRepository) Save(id string, orderData []byte) error {
	// implement cache and db save method
	return nil
}
