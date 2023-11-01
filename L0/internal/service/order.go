package service

import "github.com/gartyom/wb-practice/L0/internal/repository"

type OrderService struct {
	repo repository.OrderRepositoryInteface
}

func NewOrderService(repo repository.OrderRepositoryInteface) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) GetById(id string) ([]byte, error) {
	data, err := s.repo.GetById(id)
	return data, err
}
