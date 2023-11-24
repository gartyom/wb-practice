package main

import (
	"github.com/gartyom/wb-practice/L2/develop/dev03/internal/app"
	"github.com/gartyom/wb-practice/L2/develop/dev03/internal/pkg/config"
	"github.com/gartyom/wb-practice/L2/develop/dev03/pkg/errs"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	cfg, err := config.Get()
	errs.Check(err)

	err = app.Run(cfg)
	errs.Check(err)
}
