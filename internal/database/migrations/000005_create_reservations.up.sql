CREATE TABLE IF NOT EXISTS reservations (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER NOT NULL,
    room_id INTEGER NOT NULL,
    start_date TEXT NOT NULL,
    end_date TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    status TEXT NOT NULL
);