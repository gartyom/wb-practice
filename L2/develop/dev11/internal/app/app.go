package app

import (
	"dev11/internal/pkg/database"
	httpHandler "dev11/internal/pkg/handler/http"
	"dev11/internal/pkg/handler/middleware/logger"
	"dev11/internal/pkg/repository"
	"dev11/internal/pkg/service"
	"net/http"
)

func Run() error {
	db := database.New()
	r := repository.New(db)
	s := service.New(r)
	h := httpHandler.New(s)

	router := http.NewServeMux()
	h.InitRoutes(router)

	handler := logger.Wrap(router)

	return http.ListenAndServe(":8000", handler)
}
