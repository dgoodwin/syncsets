package api

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	_ "github.com/lib/pq"
)

var _ APIResource = &Cluster{}

// Cluster is a representation of a Cluster we will reconcile SyncSets to.
//
// swagger:model cluster
type Cluster struct {
	// Name of the cluster.
	//
	// required: true
	Name string `json:"name"`
	// Namespace of the cluster. Models the Kubernetes concept of Namespace as OpenShift Hive
	// uses that to allow multiple clusters with the same name, separated by owner.
	//
	// required: true
	Namespace string `json:"namespace"`
	// Kubeconfig is an admin kubeconfig file for communicating with the cluster.
	//
	// required: true
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

func (a Cluster) Marshal() ([]byte, error) {
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

func (a *Cluster) RowScan(row *sql.Row) error {
	return row.Scan(a)
}

func (a *Cluster) APIVersion() string {
	return "clusters"
}

type ClusterItem struct {
	ID      int
	Cluster Cluster
}
