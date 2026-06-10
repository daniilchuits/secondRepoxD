package handlers

import "database/sql"

type dbManager struct {
	db *sql.DB
}

func NewDB(db *sql.DB) dbManager {
	return dbManager{db: db}
}
