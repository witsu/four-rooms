package reservation

import (
	"testing"
	"time"

	"four-rooms/internal/database"
	dt "four-rooms/internal/datetime"
	"four-rooms/internal/inventory"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(Reservation{
		RoomID:    1,
		StartDate: parseDate("2021-01-01"),
		EndDate:   parseDate("2021-01-02"),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	err = validate.Struct(Reservation{
		RoomID:    1,
		StartDate: parseDate("2021-01-03"),
		EndDate:   parseDate("2021-01-02"),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	})
	if err == nil {
		t.Error("expected error as start date is after end date")
	}
}

func TestCreateWhenBooked(t *testing.T) {
	db := database.NewTest()
	defer db.Close()

	db.Conn().Exec("DELETE FROM room_inventory")

	if err := inventory.InsertRoomInventory(db.Conn(), 1, "2025-01-01", 1, 1); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	reserv := Reservation{
		HotelID:   1,
		RoomID:    1,
		StartDate: parseDate("2025-01-01"),
		EndDate:   parseDate("2025-01-02"),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	err := Create(db.Conn(), &reserv)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}

func TestCreateWhenOK(t *testing.T) {
	db := database.NewTest()
	defer db.Close()

	db.Conn().Exec("DELETE FROM room_inventory")
	db.Conn().Exec("DELETE FROM reservations")

	if err := inventory.InsertRoomInventory(db.Conn(), 1, "2025-01-01", 1, 0); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	reserv := Reservation{
		HotelID:   1,
		RoomID:    1,
		StartDate: parseDate("2025-01-01"),
		EndDate:   parseDate("2025-01-02"),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	err := Create(db.Conn(), &reserv)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if reserv.ID == 0 {
		t.Error("expected reservation ID > 0")
	}
	dbReserv, err := Get(db.Conn(), reserv.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if dbReserv.ID == 0 {
		t.Errorf("expected reservation, got %v", dbReserv)
	}

	db.Conn().Exec("DELETE FROM room_inventory")
	db.Conn().Exec("DELETE FROM reservations")
}

func parseDate(dateStr string) dt.Date {
	date, _ := time.Parse("2006-01-02", dateStr)
	return dt.Date(date)
}
