package reservation

import "time"

type Reservation struct {
	ID        int       `json:"id"`
	HotelID   int       `json:"hotel_id"`
	RoomID    int       `json:"room_id"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required,gtefield=StartDate"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Status    string    `json:"status"`
}
