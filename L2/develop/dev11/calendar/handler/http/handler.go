package http

import "net/http"

type CalendarService interface {
	EventsForDay(userID string, date string) error
}

type Handler struct {
	service CalendarService
}

func New(service CalendarService, router *http.ServeMux) {
	h := Handler{
		service: service,
	}

	router.HandleFunc("/events_for_day", h.eventsForDay)
	router.HandleFunc("/events_for_week", h.eventsForWeek)
	router.HandleFunc("/events_for_month", h.eventsForMonth)
	router.HandleFunc("/create_event", h.createEvent)
	router.HandleFunc("/update_event", h.updateEvent)
	router.HandleFunc("/delete_event", h.deleteEvent)
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	if userID == "" || date == "" {
		w.WriteHeader(400)
		return
	}

	err := h.service.EventsForDay(userID, date)
	if err != nil {
		w.WriteHeader(503)
	}

	w.Write([]byte(err.Error()))
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {

}
