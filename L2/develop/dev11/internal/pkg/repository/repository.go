package repository

import (
	"database/sql"
	"dev11/internal/pkg/model"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) EventsForDay(date string) ([]model.Event, error) {
	rows, err := r.db.Query("SELECT * FROM events WHERE events.date = $1", date)
	if err != nil {
		return nil, err
	}

	aE := make([]model.Event, 0)
	for rows.Next() {
		e := model.Event{}
		rows.Scan(&e.Id, &e.Date, &e.Description)
		aE = append(aE, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return aE, nil
}

func (r *Repository) EventsBetween(date1 string, date2 string) ([]model.Event, error) {
	rows, err := r.db.Query("SELECT * FROM events WHERE events.date >= $1 AND events.date <= $2", date1, date2)
	if err != nil {
		return nil, err
	}

	aE := make([]model.Event, 0)
	for rows.Next() {
		e := model.Event{}
		rows.Scan(&e.Id, &e.Date, &e.Description)
		aE = append(aE, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return aE, nil
}

func (r *Repository) CreateEvent(event model.Event) error {
	_, err := r.db.Exec("INSERT INTO events (date, description) VALUES ($1, $2)", event.Date, event.Description)
	return err
}

func (r *Repository) UpdateEvent(event model.Event) error {
	_, err := r.db.Exec("UPDATE events SET date=$1, description=$2 WHERE id=$3", event.Date, event.Description, event.Id)
	return err
}

func (r *Repository) DeleteEvent(id int) error {
	_, err := r.db.Exec("DELETE FROM events WHERE id=$1", id)
	return err
}
