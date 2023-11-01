package controller

import (
	"net/http"
	"strings"

	"github.com/gartyom/wb-practice/L0/internal/service"
)

func NewOrderController(service service.OrderServiceInterface) {
	cnt := &OrderController{
		service: service,
	}

	http.HandleFunc("/order/", cnt.OrderByUID)
}

type OrderController struct {
	service service.OrderServiceInterface
}

func (cnt *OrderController) OrderByUID(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	id := p[2]

	switch r.Method {
	case http.MethodGet:
		data, err := cnt.service.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
