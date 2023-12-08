package main

import (
	"dev06/internal/app"
	"dev06/internal/pkg/args"
	"log"
	"os"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	args, err := args.New()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	app.Run(args)
}
