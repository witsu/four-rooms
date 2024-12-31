package inventory

import "database/sql"

// Currently used for testing
// Could be used for creating room
func InsertRoom(db *sql.DB, hotelID, roomID, price int) error {
	query := `INSERT INTO rooms (id, hotel_id, price, description, size, title, type) 
		VALUES (?, ?, ?, 'desc', 100, 'Room 1', 'single')
	`
	_, err := db.Exec(query, roomID, hotelID, price)
	return err
}

// Currently used for testing
// Could be used for inserting room inventory data on daily bases
func InsertRoomInventory(db *sql.DB, roomID int, date string, total, booked int) error {
	query := `INSERT INTO room_inventory (room_id, date, total, total_booked) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, roomID, date, total, booked)
	return err
}
