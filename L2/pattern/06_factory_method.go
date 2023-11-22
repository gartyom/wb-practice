package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Фабричный метод - это порождающий паттерн проектирования, который определяет
// общий интерфейс для создания объектов в суперклассе, позволяя подклассам
// изменять тип создаваемых объектов

type Config interface {
	configurate()
}

type productionConfig struct {
	db   string
	port string
	host string
}

func NewProdConfig() Config {
	return &productionConfig{
		db:   "postgres_prod",
		port: "8080",
		host: "example.com",
	}

}

func (conf *productionConfig) configurate() {
	fmt.Println("using prod configuration:")
	fmt.Println(conf)
}

type devConfig struct {
	db   string
	port string
	host string
}

func NewDevConfig() Config {
	return &devConfig{
		db:   "postgres_dev",
		port: "8080",
		host: "localhost",
	}

}
func (conf *devConfig) configurate() {
	fmt.Println("using dev configuration")
	fmt.Println(conf)
}

func NewConfig(t string) Config { // Фабричный метод
	switch t {
	case "prod":
		return NewProdConfig()
	case "dev":
		return NewDevConfig()
	}

	return nil
}

func factoryMethodExample() {
	prodCnf := NewConfig("prod")
	devCnf := NewConfig("dev")

	prodCnf.configurate()
	devCnf.configurate()
}

// +
// Выделяет код производства продуктов в одно место, упрощая поддержку кода.
// Упрощает добавление новых продуктов в программу.
