package server

import (
	"encoding/json"
	"errors"
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
