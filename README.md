# Project blueprint-echo-sqlite-react

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

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