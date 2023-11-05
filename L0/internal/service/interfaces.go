package service

type (
	OrderServiceInterface interface {
		Save(id string, orderData []byte) error
		GetById(id string) ([]byte, error)
		HandleNewOrder(data []byte, uid string) error
		Recover() error
	}
)
