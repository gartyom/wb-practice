package app

import (
	"net/http"

	"github.com/gartyom/wb-practice/L0/internal/cacher"
	"github.com/gartyom/wb-practice/L0/internal/config"
	"github.com/gartyom/wb-practice/L0/internal/controller"
	"github.com/gartyom/wb-practice/L0/internal/repository"
	"github.com/gartyom/wb-practice/L0/internal/service"
	"github.com/gartyom/wb-practice/L0/internal/subscriber"
	"github.com/gartyom/wb-practice/L0/pkg/database"
)

func Run() error {
	cfg := config.Get()
	db := database.Connect(cfg)
	cchr := cacher.New()
	repo := repository.New(db)
	serv := service.New(repo, cchr)
	controller.New(serv)

	err := serv.Order.Recover()
	if err != nil {
		return err
	}

	sub := subscriber.New(cfg.StanClusterName, serv.Order)
	sub.Init()

	http.ListenAndServe("localhost:8000", nil)
	return nil
}
