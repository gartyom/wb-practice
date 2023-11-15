package task21

import "fmt"

func Run() {
	fmt.Println()
	fmt.Println("Task 21:")
	c := &client{}

	adapter := &postgresAdapter{
		db: &Postgres{},
	}

	c.getAll(adapter)
}

// Клиент
type client struct {
}

func (c *client) getAll(db dbInterface) {
	fmt.Println("Client")
	db.GetAll()
}

// Интерфейс бд
type dbInterface interface {
	GetAll()
}

// Адаптер
type postgresAdapter struct {
	db *Postgres
}

func (pa *postgresAdapter) GetAll() {
	fmt.Println("Adapter")
	pa.db.SelectAll()
}

// база данных
type Postgres struct {
}

func (p *Postgres) SelectAll() {
	fmt.Println("Postgres")
}
