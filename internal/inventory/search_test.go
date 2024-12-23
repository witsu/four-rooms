package inventory

import (
	"four-rooms/internal/database"
	"testing"
)

func TestSearch(t *testing.T) {
	db := database.NewTest()
	defer db.Close()

	db.Conn().Exec("DELETE FROM room_inventory")

	insertRoomInventory(t, db, 1, "2025-01-01", 1, 0)
	insertRoomInventory(t, db, 1, "2025-01-02", 1, 0)
	insertRoomInventory(t, db, 1, "2025-01-03", 1, 1)

	// Test when wrong location
	sq, err := NewSearchQuery("2025-01-01", "2025-01-02", "wrong-location")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	rooms, err := Search(db.Conn(), sq)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(rooms) != 0 {
		t.Errorf("expected 0 rooms, got %d", len(rooms))
	}

	// Test when no rooms available because 2025-01-03 is booked
	sq, err = NewSearchQuery("2025-01-01", "2025-01-03", "Los Angeles")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	rooms, err = Search(db.Conn(), sq)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(rooms) != 0 {
		t.Errorf("expected 0 rooms, got %d", len(rooms))
	}

	// Test when available rooms
	sq, err = NewSearchQuery("2025-01-01", "2025-01-02", "Los Angeles")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	rooms, err = Search(db.Conn(), sq)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(rooms) != 1 {
		t.Errorf("expected 1 room, got %d", len(rooms))
	}

	db.Conn().Exec("DELETE FROM room_inventory")
}

func insertRoomInventory(t *testing.T, db database.Service, roomID int, date string, total, booked int) {
	query := `INSERT INTO room_inventory (room_id, date, total, total_booked) VALUES (?, ?, ?, ?)`
	_, err := db.Conn().Exec(query, roomID, date, total, booked)
	if err != nil {
		t.Error(err)
	}
}
