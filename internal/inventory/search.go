package inventory

import (
	"database/sql"
	dt "four-rooms/internal/datetime"
)

type SearchQuery struct {
	StartDate dt.Date `query:"start_date" validate:"required"`
	EndDate   dt.Date `query:"end_date" validate:"required,gtefield=StartDate"`
	Location  string  `query:"location" validate:"required"`
}

func NewSearchQuery(start, end, location string) (SearchQuery, error) {
	sd, err := dt.Parse(start)
	if err != nil {
		return SearchQuery{}, err
	}
	ed, err := dt.Parse(end)
	if err != nil {
		return SearchQuery{}, err
	}
	return SearchQuery{
		StartDate: sd,
		EndDate:   ed,
		Location:  location,
	}, nil
}

func Search(db *sql.DB, sq SearchQuery) ([]Room, error) {
	query := `
		SELECT r.* FROM rooms AS r
		JOIN hotels AS h ON r.hotel_id = h.id
		JOIN room_inventory AS ri ON r.id = ri.room_id
		WHERE h.location = ?
			AND ri.date >= ? AND ri.date <= ? 
			AND ri.total > ri.total_booked 
		GROUP BY r.id
		HAVING COUNT(r.id) = (julianday(?) - julianday(?) + 1)
	`

	rows, err := db.Query(query,
		sq.Location,
		sq.StartDate.String(),
		sq.EndDate.String(),
		sq.EndDate.String(),
		sq.StartDate.String(),
	)
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
