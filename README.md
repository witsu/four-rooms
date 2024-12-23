# Four Rooms

This is a hobby project aimed at learning by creating a working hotel reservation system. The project also serves as a playground for trying out new Go and React packages.

## System Design

The design of this project is inspired by Chapter 7, "Hotel Reservation," from the book "System Design Interview - An Insider's Guide, Volume 2."

Database tables
- hotels
- rooms
- room_inventory
- reservations

API routes
- GET "/hotels",
- GET "/hotels/:id",
- GET "/hotels/:id/rooms",
- POST "/reservations",
- GET "/search",

## Project Goals

1. **Learning**: Gain hands-on experience by building a full-stack application.
2. **Experimentation**: Try out new Go packages and technologies.

## Status

**Work in Progress**

### TODO Frontend
- Implement functionality to:
  - Search available rooms by location and dates
  - Reserve a room
  - View reservation details by ID
- Improve layout by using a CSS framework, such as:
  - Beer.css
  - Bootstrap
  - Tailwind CSS

### TODO Backend
- Implement pricing:
  - Start with static pricing
  - Later, add dynamic pricing based on demand and availability
- Implement user authentication:
  - User registration
  - User login
- Implement management panel to update hotels and rooms
- Enhance error handling and validation
- Add more unit and integration tests

## Technologies Used

- **Backend**: Go, Echo, Validator
- **Database**: SQLite, Migrate
- **Frontend**: Typescript, React with router

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## DB migrations
[golang-migrate](https://github.com/golang-migrate/migrate) is used to run database migrations

Create table
```bash
migrate create -ext sql -dir internal/database/migrations -seq create_hotels_table
```

Run migration
```bash
migrate -database internal/database/hotel.db -path internal/database/migrations up
```