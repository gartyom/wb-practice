package main

import (
	"dev09/internal/app"
	"fmt"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком
*/

func main() {
	err := app.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
