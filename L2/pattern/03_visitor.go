package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Посетитель позволяет добавить функционал к классам без надобности менять код
// этих классов
type MusicVisitor interface {
	VisitRock()
	VisitPop()
	VisitRap()
}

type musicCounter struct {
	counter map[string]int
}

func (v *musicCounter) VisitRock() {
	v.counter["rock"] += 1
}
func (v *musicCounter) VisitRap() {
	v.counter["rap"] += 1
}
func (v *musicCounter) VisitPop() {
	v.counter["pop"] += 1
}

type Music interface {
	accept(v MusicVisitor)
}

type rock struct {
}

func (r *rock) accept(v MusicVisitor) {
	v.VisitRock()
}

type pop struct {
}

func (r *pop) accept(v MusicVisitor) {
	v.VisitPop()
}

type rap struct {
}

func (r *rap) accept(v MusicVisitor) {
	v.VisitRap()
}

func visitorExample() {
	musicLib := []Music{new(rock), new(rock), new(pop), new(rap)}

	v := &musicCounter{
		counter: make(map[string]int),
	}

	for _, x := range musicLib {
		x.accept(v)
	}

	fmt.Println(v.counter)
}

//+
//Упрощает добавление операций

// -

// Для добаления новой функциональноысти необходимо создавать нового visitorа
// из-за этого cложнее читать код т.к. логика одного объекта может
//находиться в нескольких visitorах
