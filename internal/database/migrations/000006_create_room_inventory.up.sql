CREATE TABLE IF NOT EXISTS room_inventory (
    room_id INTEGER NOT NULL,
    date TEXT NOT NULL,
    total INTEGER NOT NULL,
    total_booked INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (room_id, date)
);