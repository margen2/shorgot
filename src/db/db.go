package db

import (
	"database/sql"

	"github.com/margen2/shorgot/src/config"

	_ "github.com/lib/pq"
)

// ConnectDB connects to the dabatase using the StringDBConnection from the environment variables
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringDBConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
