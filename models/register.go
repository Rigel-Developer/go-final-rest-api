package models

import "github/rigel-developer/go-final-rest-api/db"

type Register struct {
	ID      int64
	UserID  int64 `binding:"required"`
	EventID int64 `binding:"required"`
}

func (r *Register) RegisterForEvent() error {
	insertRegisterSQL := `
	INSERT INTO registrations (user_id, event_id)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(insertRegisterSQL)
	if err != nil {
		return err

	}
	defer stmt.Close()

	result, err := stmt.Exec(r.UserID, r.EventID)

	if err != nil {
		return err
	}

	r.ID, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func CancelRegistration(userID, eventID int64) error {
	deleteRegisterSQL := `
	DELETE FROM registrations WHERE user_id = ? AND event_id = ?
	`
	stmt, err := db.DB.Prepare(deleteRegisterSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, eventID)
	if err != nil {
		return err
	}

	return nil
}

func GetRegistrations(userID int64) ([]Event, error) {
	queryGetRegistrations := `
	SELECT e.id, e.name, e.description, e.location, e.date_time, e.user_id
	FROM events e
	JOIN registrations r ON e.id = r.event_id
	WHERE r.user_id = ?
	`
	rows, err := db.DB.Query(queryGetRegistrations, userID)
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
