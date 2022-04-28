package repository

import (
	"database/sql"
)

type SQLStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}
