package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"four-rooms/internal/database"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func TestHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	s := &Server{}
	// Assertions
	if err := s.HelloWorldHandler(c); err != nil {
		t.Errorf("handler() error = %v", err)
		return
	}
	if resp.Code != http.StatusOK {
		t.Errorf("handler() wrong status code = %v", resp.Code)
		return
	}
	expected := map[string]string{"message": "Hello World"}
	var actual map[string]string
	// Decode the response body into the actual map
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		t.Errorf("handler() error decoding response body: %v", err)
		return
	}
	// Compare the decoded response with the expected value
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler() wrong response body. expected = %v, actual = %v", expected, actual)
		return
	}
}

func TestCreateReservationWhenMissingRequired(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}

	req := httptest.NewRequest(http.MethodPost, "/reservations", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	s := &Server{}
	// Assertions
	err := s.createReservation(c)
	if err == nil {
		t.Error("createReservation() expects error")
		return
	}
	// Check if the error is of type *echo.HTTPError
	var httpErr *echo.HTTPError
	if !errors.As(err, &httpErr) {
		t.Errorf("expected error of type *echo.HTTPError, got %T", err)
		return
	}

	if httpErr.Code != http.StatusBadRequest {
		t.Errorf("createReservation() wrong status code = %v, expected = %v", httpErr.Code, http.StatusBadRequest)
		return
	}
}

func TestCreateReservationWhenOK(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}

	payload := map[string]interface{}{
		"hotel_id":   1,
		"room_id":    1,
		"start_date": "2021-01-01",
		"end_date":   "2021-01-02",
		"first_name": "John",
		"last_name":  "Doe",
		"email":      "john.doe@example.com",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
		return
	}
	bodyReader := bytes.NewReader(jsonData)

	req := httptest.NewRequest(http.MethodPost, "/reservations", bodyReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set Content-Type header
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	s := &Server{
		db: database.NewTest(),
	}
	defer func() {
		s.db.Conn().Exec("DELETE FROM reservations")
		s.db.Close()
	}()

	// Assertions
	if err := s.createReservation(c); err != nil {
		t.Error("createReservation() doesn't expect error")
		return
	}
	if resp.Code != http.StatusOK {
		t.Errorf("handler() wrong status code = %v", resp.Code)
		return
	}
}
