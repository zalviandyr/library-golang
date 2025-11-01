# Library BE

Backend service for managing a simple library catalog. It is written in Go, uses GORM for ORM, and provides seeders to quickly populate authors, publishers, and books with realistic fake data.

## Tech Stack

- Go 1.21+
- PostgreSQL
- GORM
- Faker
- Dot.ENV

## Project Layout

- `app/domain/dao` &mdash; GORM data models (`Author`, `Publisher`, `Book`)
- `config` &mdash; environment loader and database bootstrapper
- `seeds` &mdash; seed entrypoint plus per-entity seeders
- `.env.example` &mdash; sample configuration

## Domain Overview

- `dao.Author` &mdash; stores author profile information and maintains a many-to-many relationship with books via `author_books`.
- `dao.Publisher` &mdash; publisher details with a one-to-many relationship to books.
- `dao.Book` &mdash; book metadata, foreign key to a publisher, and many-to-many authors.

## Configuration

1. Duplicate `.env.example` to `.env`.
2. Fill in database credentials (`DB_HOST`, `DB_PORT`, `DB_DATABASE`, `DB_USER`, `DB_PASSWORD`).

## Database & Seeds

Run the seeder entrypoint to populate authors and publishers (extend it with books when ready):

```bash
go run seeds/main.go
```

## Running the Application

```bash
go mod download
go run main.go
```

## Docker

Build and run the containerized binary if you prefer not to install Go locally:

```bash
docker build -t library-be .
```

The container only contains the compiled service binary, so make sure the database specified in `.env` is reachable from inside the container (e.g., use host networking or expose the database service).

## Testing

1. Create `.env.testing` (copy from `.env.example` or `.env`) and point it at an isolated PostgreSQL database that tests are free to reset.
2. Run the suite serially so that each package gets an exclusive connection:

   ```bash
   go test -p 1 ./...
   ```

   The `-p 1` flag disables parallel package execution. Each package test initializes a fresh database session and drops/migrates tables; running one at a time prevents race conditions where concurrent packages remove each other’s schema.
