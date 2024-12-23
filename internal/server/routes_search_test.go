package server

import (
	"errors"
	"four-rooms/internal/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func TestSearchWhenMissingParams(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/search", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	s := &Server{}
	// Assertions
	err := s.search(c)
	if err == nil {
		t.Error("search() expects error")
		return
	}
	// Check if the error is of type *echo.HTTPError
	var httpErr *echo.HTTPError
	if !errors.As(err, &httpErr) {
		t.Errorf("expected error of type *echo.HTTPError, got %T", err)
		return
	}

	if httpErr.Code != http.StatusBadRequest {
		t.Errorf("search() wrong status code = %v, expected = %v", httpErr.Code, http.StatusBadRequest)
		return
	}
}

func TestSearchWhenOK(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}

	url := "/search?start=2025-01-01&end=2025-01-02&location=Los%20Angeles"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	s := &Server{
		db: database.NewTest(),
	}
	defer s.db.Close()

	// Assertions
	if err := s.search(c); err != nil {
		t.Error("search() doesn't expect error")
		return
	}
	if resp.Code != http.StatusOK {
		t.Errorf("search() wrong status code = %v", resp.Code)
		return
	}
}
