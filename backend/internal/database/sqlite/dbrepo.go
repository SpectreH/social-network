package sqlite

import (
	"database/sql"
	"log"
	"social-network/internal/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// sqliteDBRepo holds the database for sqlite
type sqliteDBRepo struct {
	DB *sql.DB
}

// SetSqliteRepo sets repository for sqlite
func SetSqliteRepo(conn *sql.DB) database.DatabaseRepo {
	return &sqliteDBRepo{
		DB: conn,
	}
}

func UseMigrations(database *sql.DB) error {
	driver, err := sqlite3.WithInstance(database, &sqlite3.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance("file://internal/database/migrations/sqlite/", "sqlite3", driver)
	if err != nil {
		return err
	}

	if err = migration.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
		log.Println("There no any changes with migration")
	}

	return nil
}
