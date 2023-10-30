package main

import (
	"fmt"
	"net/http"

	"github.com/gartyom/wb-pratctice/L0/config"
	"github.com/gartyom/wb-pratctice/L0/database"
	"github.com/nats-io/stan.go"
)

func main() {
	cfg := config.Get()

	sc, _ := stan.Connect(cfg.StanClusterName, "sub-1")

	database.Connect(cfg)

	sc.Subscribe("orders", func(m *stan.Msg) {
		fmt.Println("message recieved: " + string(m.Data))
	})

	http.ListenAndServe("localhost:8000", nil)
}
