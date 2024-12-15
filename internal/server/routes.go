package server

import (
	"four-rooms/internal/inventory"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/", s.HelloWorldHandler)
	e.GET("/health", s.healthHandler)

	e.GET("/hotels", s.getHotels)
	e.GET("/hotels/:id", s.getHotel)

	return e
}

func (s *Server) getHotels(c echo.Context) error {
	hotels, err := inventory.GetHotels(s.db.Conn())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, hotels)
}

func (s *Server) getHotel(c echo.Context) error {
	id := c.Param("id")
	hotel, err := inventory.GetHotel(s.db.Conn(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "hotel not found")
	}
	return c.JSON(http.StatusOK, hotel)
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
