package api

import (
	"database/sql"
)

type APIResource interface {
	RowScan(row *sql.Row) error
	Marshal() ([]byte, error)
	Scan(value interface{}) error
	APIVersion() string // to find db tables and api endpoints
}
