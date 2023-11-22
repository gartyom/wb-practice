package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «строитель».
    Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Строитель - порождающий паттерн проектирования, который позволяет создавать объекты пошагово

type AbstractBuilder interface {
	buildPartA()
	buildPartB()
	GetObject() *obj
}

type obj struct {
	objHeight int
	objWidth  int
}

type builder1 struct {
	o obj
}

func (b *builder1) buildPartA() {
	b.o.objHeight = 1
}

func (b *builder1) buildPartB() {
	b.o.objWidth = 2
}

func (b *builder1) GetObject() *obj {
	return &b.o
}

type director struct {
	b AbstractBuilder
}

func (d *director) Construct(b AbstractBuilder) *obj {
	b.buildPartA()
	b.buildPartB()
	return b.GetObject()
}

func build() {
	var b builder1

	d := director{}

	_ = d.Construct(&b)
}

// strings.bulder является примером паттерна builder, но без директора

func test() {
	var s strings.Builder
	s.WriteByte(byte(0))
	s.WriteString("a")
	s.WriteRune([]rune("a")[0])
	_ = s.String()

	var b builder1
	b.buildPartA()
	b.buildPartB()
	_ = b.GetObject()
}

// +
// позволяет менять "внутренности" создаваемых объектов
// Разделение сборки объекта от его бизнес логики
// Дает контроль на каждом этапе сборки

// -
// Отдельный builder для разных объектов
//
