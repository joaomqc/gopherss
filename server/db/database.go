package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	migrate "github.com/golang-migrate/migrate"
	sqlite3 "github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
)

type DB struct {
	*sql.DB
}

var _db *DB

func GetDb() (*DB, error) {
	if _db == nil {
		sqlDb, err := sql.Open("sqlite3", "gopherss.db")
		if err != nil {
			return nil, err
		}
		_db = &DB{
			DB: sqlDb,
		}
	}
	return _db, nil
}

func (db *DB) Initialize() error {
	driver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
