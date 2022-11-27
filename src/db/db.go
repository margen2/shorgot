package db

import (
	"github.com/margen2/shorgot/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
// db, err = sql.Open("postgres", connStr)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringDBConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
