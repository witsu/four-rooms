package reservation

import (
	"testing"
	"time"

	dt "four-rooms/internal/datetime"

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

func parseDate(dateStr string) dt.Date {
	date, _ := time.Parse("2006-01-02", dateStr)
	return dt.Date(date)
}
