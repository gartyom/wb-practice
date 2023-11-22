package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Состояние — это поведенческий паттерн проектирования, который позволяет
// объектам менять поведение в зависимости от своего состояния. Извне создаётся
// впечатление, что изменился класс объекта.

type State interface {
	getRoot()
	getRegularData()
	getSecrets()
}

type user struct {
	noroot State
	root   State

	currentState State
}

func (u *user) requestRootAccess() {
	u.currentState.getRoot()
}
func (u *user) requestRegularData() {
	u.currentState.getRegularData()
}
func (u *user) requestSecretData() {
	u.currentState.getSecrets()
}

type norootState struct {
	user *user
}

func (nr *norootState) getRoot() {
	nr.user.currentState = nr.user.root
	fmt.Println("root granted")
}
func (nr *norootState) getRegularData() {
	fmt.Println("regular data")
}
func (nr *norootState) getSecrets() {
	fmt.Println("fail: non root user")
}

type rootState struct {
	user *user
}

func (r *rootState) getRoot() {
	fmt.Println("root already granted")
}
func (nr *rootState) getRegularData() {
	fmt.Println("regular data")
}
func (nr *rootState) getSecrets() {
	fmt.Println("secrets")
}

func stateExample() {
	u := &user{}

	noroot := &norootState{user: u}
	root := &rootState{user: u}

	u.noroot = noroot
	u.root = root
	u.currentState = noroot

	u.requestRegularData()
	u.requestSecretData()
	u.requestRootAccess()
	u.requestSecretData()

}

// +
// Избавляет от множества больших условных операторов машины состояний.
// Концентрирует в одном месте код, связанный с определённым состоянием.

// -
// Может неоправданно усложнить код, если состояний мало и они редко меняются.
