package reservation

import (
	"database/sql"
	dt "four-rooms/internal/datetime"
)

const STATUS_PENDING = "pending"
const STATUS_CANCELLED = "cancelled"
const STATUS_PAID = "paid"

type Reservation struct {
	ID        int     `json:"id"`
	HotelID   int     `json:"hotel_id"`
	RoomID    int     `json:"room_id" validate:"required"`
	StartDate dt.Date `json:"start_date" validate:"required"`
	EndDate   dt.Date `json:"end_date" validate:"required,gtefield=StartDate"`
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	Email     string  `json:"email" validate:"required,email"`
	Status    string  `json:"status"`
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
