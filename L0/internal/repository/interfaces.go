package repository

type (
	OrderRepositoryInteface interface {
		GetById(id string) ([]byte, error)
	}
)
