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
	Price       int    `json:"price"`
}

func GetHotels(db *sql.DB) ([]Hotel, error) {
	rows, err := db.Query("SELECT id, name, address, location FROM hotels")
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
	row := db.QueryRow("SELECT id, name, address, location FROM hotels WHERE id = ?", id)

	var hotel Hotel
	err := row.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.Location)
	if err != nil {
		return Hotel{}, err
	}

	return hotel, nil
}

func GetHotelRooms(db *sql.DB, id string) ([]Room, error) {
	query := `
		SELECT id, hotel_id, description, size, title, type, price
		FROM rooms 
		WHERE hotel_id = ?
	`
	rows, err := db.Query(query, id)
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
			&room.Price,
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

func GetHotelRoom(db *sql.DB, hotelID, roomID int) (Room, error) {
	query := `
		SELECT id, hotel_id, description, size, title, type, price 
		FROM rooms 
		WHERE hotel_id = ? AND id = ?
	`
	row := db.QueryRow(query, hotelID, roomID)

	var room Room
	err := row.Scan(
		&room.ID,
		&room.HotelID,
		&room.Description,
		&room.Size,
		&room.Title,
		&room.Type,
		&room.Price,
	)
	if err != nil {
		return Room{}, err
	}

	return room, nil
}
