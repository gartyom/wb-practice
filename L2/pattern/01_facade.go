package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
    Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Фасад предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку

type service1 struct {
}

func (s *service1) Save(name string) {
	fmt.Println("saved: ", name)
	return
}

type service2 struct {
}

func (s *service2) Notify() {
	fmt.Println("notified")
	return
}

type facade struct {
	s1 service1
	s2 service2
}

func (f *facade) AddNew(name string) {
	f.s1.Save(name)
	f.s2.Notify()
}

func testFacade() {
	var f facade
	f.AddNew("asd")
}

//при созданиий http сервера по архитектуре model service controller, controller выступает в роли фасада

//+
//Изолирует пользователей системы от сложных компонент

//-
//Может получится так, что в конечном счете фасад будет хранить и делать сдишком много
