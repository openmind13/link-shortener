package store

import "database/sql"

// Store struct
type Store struct {
	db *sql.DB
}

// New - return new store
func New() *Store {
	return &Store{}
}
