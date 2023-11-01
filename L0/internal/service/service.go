package service

import "github.com/gartyom/wb-practice/L0/internal/repository"

type Service struct {
	Order OrderServiceInterface
}

func New(r *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(r.Order),
	}
}
