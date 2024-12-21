package reservation

import (
	"database/sql"
	"time"
)

const STATUS_PENDING = "pending"
const STATUS_CANCELLED = "cancelled"
const STATUS_PAID = "paid"

type Reservation struct {
	ID        int       `json:"id"`
	HotelID   int       `json:"hotel_id"`
	RoomID    int       `json:"room_id" validate:"required"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required,gtefield=StartDate"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Status    string    `json:"status"`
}

func Create(db *sql.DB, reserv *Reservation) error {
	query := `
		INSERT INTO reservations 
		(hotel_id, room_id, start_date, end_date, first_name, last_name, email, status)
		VALUES (?,?,?,?,?,?,?,?)
		ON DUPLICATE KEY UPDATE
			users = VALUES(users)
	`
	_, err := db.Exec(query,
		reserv.HotelID,
		reserv.RoomID,
		reserv.StartDate.Format(time.DateOnly),
		reserv.EndDate.Format(time.DateOnly),
		reserv.FirstName,
		reserv.LastName,
		reserv.Email,
		STATUS_PENDING)
	return err
}
