CREATE TABLE IF NOT EXISTS rooms (
    id INTEGER PRIMARY KEY,
    hotel_id INTEGER NOT NULL,
    description TEXT,
    size INTEGER NOT NULL,
    title TEXT NOT NULL,
    type TEXT NOT NULL
);