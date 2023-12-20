package service

import (
	"dev11/internal/pkg/model"
	"dev11/pkg/date"
	"encoding/json"
)

type Repository interface {
	EventsBetween(date1 string, date2 string) ([]model.Event, error)
	EventsForDay(date string) ([]model.Event, error)
	CreateEvent(event model.Event) error
	UpdateEvent(event model.Event) error
	DeleteEvent(id int) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) EventsForDay(date string) ([]byte, error) {
	events, err := s.repo.EventsForDay(date)
	if err != nil {
		return nil, err
	}

	encoded, err := json.Marshal(events)
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

func (s *Service) EventsForWeek(d string) ([]byte, error) {
	return s.EventsForDay(d)
}

func (s *Service) EventsForMonth(d string) ([]byte, error) {
	d1, d2, err := date.GetMonthPeriod(d)
	if err != nil {
		return nil, err
	}

	return s.EventsBetween(d1, d2)
}

func (s *Service) EventsBetween(date1 string, date2 string) ([]byte, error) {
	if date1 > date2 {
		date1, date2 = date2, date1
	}

	events, err := s.repo.EventsBetween(date1, date2)
	if err != nil {
		return nil, err
	}

	encoded, err := json.Marshal(events)
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

func (s *Service) CreateEvent(event model.Event) error {
	err := s.repo.CreateEvent(event)
	return err
}

func (s *Service) UpdateEvent(event model.Event) error {
	err := s.repo.UpdateEvent(event)
	return err
}

func (s *Service) DeleteEvent(id int) error {
	err := s.repo.DeleteEvent(id)
	return err
}
