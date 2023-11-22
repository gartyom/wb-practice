package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Цепочка обязанностей - поведенческий паттерн проектирования, который
// позволяет передавать запросы последовательно по цепочке обработчиков. Каждый
// последующий обработчик решает, может ли он обработать запрос сам и стоит ли
// передавать запрос дальше по цепи.

type data struct {
	validated bool
	saved     bool
	notified  bool
}

type Handler interface {
	setNext(h Handler)
	process(d *data)
}

type validator struct {
	next Handler
}

func (v *validator) setNext(h Handler) {
	v.next = h

}
func (v *validator) process(d *data) {
	if d.validated {
		fmt.Println("data already validated")
		v.next.process(d)
	}

	fmt.Println("data successfully validated")
	d.validated = true
	v.next.process(d)

}

type chache struct {
	next Handler
}

func (c *chache) setNext(h Handler) {
	c.next = h
}
func (c *chache) process(d *data) {
	if d.saved {
		fmt.Println("data already saved")
		c.next.process(d)
	}

	fmt.Println("data successfully saved")
	d.saved = true
	c.next.process(d)

}

type notificator struct {
	next Handler
}

func (n *notificator) setNext(h Handler) {
	n.next = h
}
func (n *notificator) process(d *data) {
	if d.notified {
		fmt.Println("service already notified")
		n.next.process(d)
	}

	fmt.Println("service succesfully notified")
	d.notified = true
	n.next.process(d)

}

type ender struct {
}

func (e *ender) setNext(h Handler) {
}

func (e *ender) process(d *data) {
	fmt.Println("chain succecfully ended")
	fmt.Println(d)
}

func chainExample() {
	v := &validator{}
	c := &chache{}
	n := &notificator{}

	n.setNext(&ender{})
	c.setNext(n)
	v.setNext(c)

	var d data

	v.process(&d)
}

// +
//Уменьшает зависимость между клиентом и обработчиками.
//Реализует принцип единственной обязанности.

// -
// нет гарантий что запрос будет обработан
