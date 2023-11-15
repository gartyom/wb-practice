package task14

import (
	"fmt"
	"reflect"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 14:")
	c := make(chan string)
	i := 0
	s := ""
	b := true

	fmt.Println(getType(c))
	fmt.Println(getType(i))
	fmt.Println(getType(s))
	fmt.Println(getType(b))
}

func getType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func getType1(v interface{}) string {
	if _, ok := v.(string); ok {
		return "string"
	}
	if _, ok := v.(chan string); ok {
		return "chan string"
	}
	if _, ok := v.(int); ok {
		return "int"
	}
	if _, ok := v.(bool); ok {
		return "bool"
	}
	return "not supported"
}
