package main

import (
	"flag"
	"fmt"
	"errors"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // Import the SQLite database driver
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "Path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "Path to migrations") // Fixed missing comma
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "Name of migration table")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
			return
		}
		panic(err)
	}
}
