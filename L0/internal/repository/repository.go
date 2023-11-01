package repository

import "github.com/gartyom/wb-practice/L0/internal/cache"

type Repository struct {
	Order OrderRepositoryInteface
}

func New(cch *cache.Cache) *Repository {
	return &Repository{
		Order: NewOrderRepository(cch),
	}
}
