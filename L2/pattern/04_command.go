package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// это поведенческий паттерн проектирования, который превращает запросы в
// объекты, позволяя передавать их как аргументы при вызове методов, ставить
// запросы в очередь, логировать их, а также поддерживать отмену операций.

type command interface {
	Do()
	Undo()
}

type database struct {
}

func (db *database) saveModel1(model string) {
	fmt.Println(model + "saved")
}
func (db *database) deleteModel1(model string) {
	fmt.Println(model + "deleted")
}
func (db *database) saveModel2(model string) {
	fmt.Println(model + "saved")
}
func (db *database) deleteModel2(model string) {
	fmt.Println(model + "deleted")
}

type saveModel1Command struct {
	db    *database
	model string
}

func (cmd *saveModel1Command) Do() {
	cmd.db.saveModel1(cmd.model)
}
func (cmd *saveModel1Command) Undo() {
	cmd.db.deleteModel1(cmd.model)
}

type saveModel2Command struct {
	db    *database
	model string
}

func (cmd *saveModel2Command) Do() {
	cmd.db.saveModel1(cmd.model)
}
func (cmd *saveModel2Command) Undo() {
	cmd.db.deleteModel1(cmd.model)
}

func commandExmaple() {
	cmdList := []command{}
	db := &database{}

	s1 := &saveModel1Command{
		db:    db,
		model: "model1",
	}

	s2 := &saveModel2Command{
		db:    db,
		model: "model2",
	}

	s1.Do()
	cmdList = append(cmdList, s1)
	s2.Do()
	cmdList = append(cmdList, s2)

	// if something unexpected happened, undo
	for _, cmd := range cmdList {
		cmd.Undo()
	}
}

// +
// Убирает прямую зависимость между объектами, вызывающими операции, и
// объектами, которые их непосредственно выполняют.

// Позволяет реализовать простую отмену и повтор операций.

// Позволяет собирать сложные команды из простых.

// -
// Усложнение кода программы
