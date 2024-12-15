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

func GetHotels(db *sql.DB) ([]Hotel, error) {
	rows, err := db.Query("SELECT * FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []Hotel
	for rows.Next() {
		var hotel Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Location); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hotels, nil
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
