package controller

import "github.com/gartyom/wb-practice/L0/internal/service"

func New(srv *service.Service) {
	NewOrderController(srv.Order)
}
