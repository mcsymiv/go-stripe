package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("unable to connect to DB. Error: %v", err)

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("unable to ping DB. Error: %v", err)

		return nil, err
	}

	return db, nil
}
