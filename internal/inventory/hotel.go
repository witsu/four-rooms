package inventory

import (
	"database/sql"
)

type Hotel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Location string `json:"location"`
}

func GetHotel(db *sql.DB, id string) (Hotel, error) {
	row := db.QueryRow("SELECT * FROM hotels WHERE id = ?", id)

	var hotel Hotel
	err := row.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.Location)
	if err != nil {
		return Hotel{}, err
	}

	return hotel, nil
}
