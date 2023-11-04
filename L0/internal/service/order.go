package service

import (
	"errors"

	"github.com/gartyom/wb-practice/L0/internal/repository"
)

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

func (s *OrderService) Save(id string, orderData []byte) error {
	//err := s.repo.Save(id, orderData)
	return nil
}

func (s *OrderService) HandleNewOrder(data []byte, uid string) error {
	_, err := s.GetById(uid)
	if err == nil {
		return errors.New("Order already exist")
	} else if err.Error() != "Cache: Not found" {
		return err
	}
	s.Save(uid, data)
	return nil
}
