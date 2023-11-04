package app

import (
	"net/http"

	"github.com/gartyom/wb-practice/L0/internal/cache"
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
	cch := cache.New()
	cch.Recover(db)
	repo := repository.New(cch)
	serv := service.New(repo)
	controller.New(serv)

	subscriber.Init(cfg.StanClusterName, serv.Order)

	http.ListenAndServe("localhost:8000", nil)
	return nil
}
