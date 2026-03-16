package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return db, nil
}

func migrate(db *sql.DB) error {
	query := `
CREATE TABLE IF NOT EXISTS telemetry (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	speed REAL NOT NULL,
	rpm REAL NOT NULL,
	temperature REAL NOT NULL,
	acceleration REAL NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
);`
	_, err := db.Exec(query)
	return err
}
