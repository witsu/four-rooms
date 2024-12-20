package reservation

type Reservation struct {
	ID        int    `json:"id"`
	HotelID   int    `json:"hotel_id"`
	RoomID    int    `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}
