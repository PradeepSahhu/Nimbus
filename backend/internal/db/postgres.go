package db

import "database/sql"

func InitDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=nimbus password=nimbus dbname=nimbus_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
