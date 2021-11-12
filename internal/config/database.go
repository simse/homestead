package config

import (
	"github.com/asdine/storm/v3"
)

// OpenDatabase opens a database connection.
func CreateDB(uri string) *storm.DB {
	db, err := storm.Open(uri)

	if err != nil {
		panic(err)
	}

	return db
}
