package psql

import (
	"database/sql"
	"fmt"
)

func Connect(uri string) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("could not open a connection to postgres: \nuri: %s\nerror: %v", uri, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping postgres: \nuri: %s\nerror: %v", uri, err)
	}

	return db, nil
}
