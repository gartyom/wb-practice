package http

import (
	"dev11/internal/pkg/model"
	"dev11/pkg/date"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service interface {
	EventsForDay(date string) ([]byte, error)
	EventsForWeek(date string) ([]byte, error)
	EventsForMonth(date string) ([]byte, error)
	CreateEvent(event model.Event) error
	UpdateEvent(event model.Event) error
	DeleteEvent(id int) error
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes(router *http.ServeMux) {
	router.HandleFunc("/events_for_day", h.eventsForDay)
	router.HandleFunc("/events_for_week", h.eventsForWeek)
	router.HandleFunc("/events_for_month", h.eventsForMonth)
	router.HandleFunc("/create_event", h.createEvent)
	router.HandleFunc("/update_event", h.updateEvent)
	router.HandleFunc("/delete_event", h.deleteEvent)
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	d := r.URL.Query().Get("date")
	if !date.Validate(d) {
		w.WriteHeader(400)
		return
	}

	data, err := h.service.EventsForDay(d)
	if err != nil {
		w.WriteHeader(503)
		return
	}

	w.Write(data)
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	d := r.URL.Query().Get("date")
	if !date.Validate(d) {
		w.WriteHeader(400)
		return
	}

	data, err := h.service.EventsForWeek(d)
	if err != nil {
		w.WriteHeader(503)
		return
	}

	w.Write(data)
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	d := r.URL.Query().Get("date")
	if !date.Validate(d) {
		w.WriteHeader(400)
		return
	}

	data, err := h.service.EventsForMonth(d)
	if err != nil {
		w.WriteHeader(503)
		return
	}

	w.Write(data)
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	var event model.Event
	json.NewDecoder(r.Body).Decode(&event)

	if event.Date == "" || event.Description == "" {
		w.WriteHeader(400)
		return
	}

	err := h.service.CreateEvent(event)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\": %s}", err.Error())))
		return
	}

	w.Write([]byte(`{"result": "event created successfully"}`))
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	var event model.Event
	json.NewDecoder(r.Body).Decode(&event)

	if event.Date == "" || event.Description == "" {
		w.WriteHeader(400)
		return
	}

	err := h.service.UpdateEvent(event)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\": %s}", err.Error())))
		return
	}

	w.Write([]byte(`{"result": "event created successfully"}`))

}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	var event model.Event
	json.NewDecoder(r.Body).Decode(&event)

	if event.Id == 0 {
		w.WriteHeader(400)
		return
	}

	err := h.service.DeleteEvent(event.Id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\": %s}", err.Error())))
		return
	}

	w.Write([]byte(`{"result": "event created successfully"}`))
}
