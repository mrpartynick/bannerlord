package migrator

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func Make() {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5433/bannerlord?sslmode=disable")
	if err != nil {
		log.Fatalf("some troubles: %w", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("some troubles: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/migrator/migrations",
		"bannerlord", driver)
	if err != nil {
		err = err.(*os.PathError)
		log.Fatal(err)
	}
	m.Up()
}
