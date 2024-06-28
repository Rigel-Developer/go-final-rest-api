package models

import (
	"github/rigel-developer/go-final-rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e *Event) Save() error {

	insertEventSQL := `
	INSERT INTO events (name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(insertEventSQL)
	if err != nil {
		return err

	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	e.ID, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func GetAll() ([]Event, error) {
	queryGetAllEvents := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(queryGetAllEvents)
	if err != nil {
		return nil, err

	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil

}

func GetOne(id int64) (Event, error) {
	queryGetOneEvent := `
	SELECT * FROM events WHERE id = ?
	`
	var e Event
	err := db.DB.QueryRow(queryGetOneEvent, id).Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return e, err
	}
	return e, nil
}

func (e *Event) Update() error {
	updateEventSQL := `
	UPDATE events SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(updateEventSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}
	return nil
}
