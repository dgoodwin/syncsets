package api

// Re-using SQL interfaces to define our json marshalling functions for now.
type APIResource interface {
	Scan(value interface{}) error
	Marshal() ([]byte, error)
	APIVersion() string // to find db tables and api endpoints
}
