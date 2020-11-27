package api

import (
	"database/sql/driver"
)

// Re-using SQL interfaces to define our json marshalling functions for now.
type APIResource interface {
	Value() (driver.Value, error)
	Scan(value interface{}) error
	APIVersion() string // to find db tables and api endpoints
}
