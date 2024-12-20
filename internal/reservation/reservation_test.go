package reservation

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(Reservation{
		StartDate: parseTime("2021-01-01"),
		EndDate:   parseTime("2021-01-02"),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	err = validate.Struct(Reservation{
		StartDate: parseTime("2021-01-03"),
		EndDate:   parseTime("2021-01-02"),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	})
	if err == nil {
		t.Error("expected error for invalid start date")
	}
}

func parseTime(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
