package inventory

import (
	"four-rooms/internal/database"
	"testing"
)

func TestSearch(t *testing.T) {
	db := database.NewTest()
	defer db.Close()

	db.Conn().Exec("DELETE FROM room_inventory")

	if err := InsertRoomInventory(db.Conn(), 1, "2025-01-01", 1, 0); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if err := InsertRoomInventory(db.Conn(), 1, "2025-01-02", 1, 1); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

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
