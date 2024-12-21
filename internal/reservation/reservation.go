package reservation

import (
	"database/sql"
	"time"
)

const STATUS_PENDING = "pending"
const STATUS_CANCELLED = "cancelled"
const STATUS_PAID = "paid"

type Reservation struct {
	ID        int    `json:"id"`
	HotelID   int    `json:"hotel_id"`
	RoomID    int    `json:"room_id" validate:"required"`
	StartDate Date   `json:"start_date" validate:"required"`
	EndDate   Date   `json:"end_date" validate:"required,gtefield=StartDate"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Status    string `json:"status"`
}

type Date time.Time

func (d *Date) UnmarshalJSON(bytes []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(date)

	return nil
}

func (d *Date) String() string {
	return time.Time(*d).Format(time.DateOnly)
}

func Create(db *sql.DB, reserv *Reservation) error {
	query := `
		INSERT INTO reservations 
		(hotel_id, room_id, start_date, end_date, first_name, last_name, email, status)
		VALUES (?,?,?,?,?,?,?,?)
	`
	_, err := db.Exec(query,
		reserv.HotelID,
		reserv.RoomID,
		reserv.StartDate.String(),
		reserv.EndDate.String(),
		reserv.FirstName,
		reserv.LastName,
		reserv.Email,
		STATUS_PENDING)
	return err
}
