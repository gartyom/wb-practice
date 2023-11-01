package service

type (
	OrderServiceInterface interface {
		GetById(id string) ([]byte, error)
	}
)
