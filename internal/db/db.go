package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type sqlDB struct {
	DB *sql.DB
}

func (s *sqlDB) Begin() {
	s.DB.Begin()
}

func initSqlDB(db *sql.DB) *sqlDB {
	return &sqlDB{
		DB: db,
	}
}
