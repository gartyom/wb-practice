package app

import (
	"fmt"
	"net/http"

	"github.com/gartyom/wb-practice/L0/internal/cache"
	"github.com/gartyom/wb-practice/L0/internal/config"
	"github.com/gartyom/wb-practice/L0/internal/controller"
	"github.com/gartyom/wb-practice/L0/internal/repository"
	"github.com/gartyom/wb-practice/L0/internal/service"
	"github.com/gartyom/wb-practice/L0/pkg/database"
	"github.com/nats-io/stan.go"
)

func Run() error {
	cfg := config.Get()

	sc, err := stan.Connect(cfg.StanClusterName, "order-reciever")

	if err != nil {
		return err
	}

	db := database.Connect(cfg)
	cch := cache.New()
	cch.Recover(db)
	repo := repository.New(cch)
	serv := service.New(repo)
	controller.New(serv)

	sc.Subscribe("orders", func(m *stan.Msg) {
		fmt.Println("message recieved: " + string(m.Data))
	})

	http.ListenAndServe("localhost:8000", nil)
	return nil
}
