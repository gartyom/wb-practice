package service

import (
	"errors"

	"github.com/gartyom/wb-practice/L0/internal/cacher"
	"github.com/gartyom/wb-practice/L0/internal/repository/postgres"
)

type OrderService struct {
	repo   postgres.PostgresOrderRepositoryInteface
	cacher cacher.CacherInterface
}

func NewOrderService(repo postgres.PostgresOrderRepositoryInteface, cacher cacher.CacherInterface) *OrderService {
	return &OrderService{
		repo:   repo,
		cacher: cacher,
	}
}

func (s *OrderService) GetById(id string) ([]byte, error) {
	data, err := s.cacher.GetById(id)
	return data, err
}

func (s *OrderService) Save(id string, orderData []byte) error {
	err := s.repo.Save(id, orderData)
	if err != nil {
		return err
	}

	s.cacher.Save(id, orderData)
	return nil
}

func (s *OrderService) HandleNewOrder(data []byte, uid string) error {
	_, err := s.GetById(uid)
	if err == nil {
		return errors.New("Order already exist")
	} else if err.Error() != "Cache: Not found" {
		return err
	}
	err = s.Save(uid, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) Recover() error {
	orederArr, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	for _, order := range orederArr {
		s.cacher.Save(order.OrderUID, order.Data)
	}

	return nil
}
