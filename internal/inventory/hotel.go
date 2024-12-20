package inventory

import (
	"database/sql"
)

type Hotel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Location string `json:"location"`
}

type Room struct {
	ID          int    `json:"id"`
	HotelID     int    `json:"hotel_id"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Title       string `json:"title"`
	Type        string `json:"type"`
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
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.Location); err != nil {
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

func GetHotelRooms(db *sql.DB, id string) ([]Room, error) {
	rows, err := db.Query("SELECT * FROM rooms WHERE hotel_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []Room
	for rows.Next() {
		var room Room
		if err := rows.Scan(
			&room.ID,
			&room.HotelID,
			&room.Description,
			&room.Size,
			&room.Title,
			&room.Type,
		); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
