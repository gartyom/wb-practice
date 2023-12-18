package app

import (
	calendarHttpHandler "dev11/calendar/handler/http"
	"dev11/calendar/handler/middleware/logger"
	"net/http"
)

func Run() error {
	router := http.NewServeMux()
	calendarHttpHandler.New(router)

	handler := logger.Wrap(router)

	return http.ListenAndServe(":8000", handler)
}
