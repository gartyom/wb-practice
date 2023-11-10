package controller

import "github.com/gartyom/wb-practice/L0/internal/service"

func HandleRequests(srv *service.Service) {
	InitOrderController(srv.Order)
}
