package api

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	_ "github.com/lib/pq"
)

type ClusterItem struct {
	ID      int
	Cluster Cluster
}

type Cluster struct {
	Name       string `json:"name"`
	Kubeconfig string `json:"kubeconfig"`
	/*
		Ingredients []string `json:"ingredients,omitempty"`
		Organic     bool     `json:"organic,omitempty"`
		Dimensions  struct {
			Weight float64 `json:"weight,omitempty"`
		} `json:"dimensions,omitempty"`
	*/
}

func (a Cluster) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implement the sql.Scanner interface to decode a JSON-encoded value into the struct fields.
func (a *Cluster) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
