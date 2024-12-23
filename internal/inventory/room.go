package inventory

import "database/sql"

// Currently used for testing
// Could be used for inserting room inventory data on daily bases
func InsertRoomInventory(db *sql.DB, roomID int, date string, total, booked int) error {
	query := `INSERT INTO room_inventory (room_id, date, total, total_booked) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, roomID, date, total, booked)
	return err
}
