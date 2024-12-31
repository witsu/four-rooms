package reservation

import (
	"database/sql"
	dt "four-rooms/internal/datetime"
)

const STATUS_PENDING = "pending"
const STATUS_CANCELLED = "cancelled"
const STATUS_PAID = "paid"

type Reservation struct {
	ID         int     `json:"id"`
	HotelID    int     `json:"hotel_id"`
	RoomID     int     `json:"room_id" validate:"required"`
	StartDate  dt.Date `json:"start_date" validate:"required"`
	EndDate    dt.Date `json:"end_date" validate:"required,gtefield=StartDate"`
	FirstName  string  `json:"first_name" validate:"required"`
	LastName   string  `json:"last_name" validate:"required"`
	Email      string  `json:"email" validate:"required,email"`
	Status     string  `json:"status"`
	TotalPrice int     `json:"total_price"`
}

func Create(db *sql.DB, reserv *Reservation, price int) error {
	reserv.Status = STATUS_PENDING
	reserv.TotalPrice = calculateTotalPrice(reserv.StartDate, reserv.EndDate, price)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	query := `
		INSERT INTO reservations 
		(hotel_id, room_id, start_date, end_date, first_name, last_name, email, status, total_price)
		VALUES (?,?,?,?,?,?,?,?,?)
	`
	res, err := tx.Exec(query,
		reserv.HotelID,
		reserv.RoomID,
		reserv.StartDate.String(),
		reserv.EndDate.String(),
		reserv.FirstName,
		reserv.LastName,
		reserv.Email,
		reserv.Status,
		reserv.TotalPrice,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `
		UPDATE room_inventory 
		SET total_booked = total_booked + 1 
		WHERE room_id = ? AND date >= ? AND date < ?
	`
	_, err = tx.Exec(updateQuery, reserv.RoomID, reserv.StartDate.String(), reserv.EndDate.String())
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err == nil {
		reserv.ID = int(id)
	}
	return err
}

// TODO add route to get reservation by id
func Get(db *sql.DB, id int) (Reservation, error) {
	query := `
		SELECT id, hotel_id, room_id, start_date, end_date, first_name, last_name, email, status, total_price
		FROM reservations 
		WHERE id = ?
	`
	row := db.QueryRow(query, id)

	var reserv Reservation
	var start, end string
	err := row.Scan(
		&reserv.ID,
		&reserv.HotelID,
		&reserv.RoomID,
		&start,
		&end,
		&reserv.FirstName,
		&reserv.LastName,
		&reserv.Email,
		&reserv.Status,
		&reserv.TotalPrice,
	)
	if err != nil {
		return Reservation{}, err
	}
	reserv.StartDate, _ = dt.Parse(start)
	reserv.EndDate, _ = dt.Parse(start)

	return reserv, nil
}

func calculateTotalPrice(start, end dt.Date, price int) int {
	return int(end.Time().Sub(start.Time()).Hours()/24) * price
}
