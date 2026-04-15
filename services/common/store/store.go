package store

import "github.com/jmoiron/sqlx"

type Store interface{}

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}
