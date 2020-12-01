package handlers

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"github.com/dgoodwin/syncsets/api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ClusterHandler struct {
	logger   log.FieldLogger
	db       *sql.DB
	registry *api.Registry
}

func NewClusterHandler(db *sql.DB) *ClusterHandler {
	return &ClusterHandler{
		db:       db,
		logger:   log.WithField("handler", "cluster"),
		registry: api.NewRegistry(),
	}
}

func GetResourceType(req *http.Request) (string, error) {
	return req.URL.Path[1:], nil
}

func (h *ClusterHandler) Get(resp http.ResponseWriter, req *http.Request) {
	h.logger.Info("called get handler")
	resource, err := GetResourceType(req)
	if err != nil {
		log.WithError(err).Error("error getting resource")
		return
	}
	h.logger.Infof("working with resource: %s", resource)
	r, err := h.registry.GetResource(resource)
	if err != nil {
		log.WithError(err).Error("error getting resource")
		return
	}
	row := h.db.QueryRow(
		fmt.Sprintf("SELECT data FROM %s ORDER BY id DESC LIMIT 1", resource))

	err = r.RowScan(row)
	if err != nil {
		log.WithError(err).Error("error querying db")
		return
	}
	jsonBytes, err := r.Marshal()
	if err != nil {
		log.WithError(err).Error("error marshalling json")
		return
	}
	fmt.Fprintf(resp, string(jsonBytes))
}

func (h *ClusterHandler) Post(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "called post handler")
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithError(err).Error("error reading request body")
		fmt.Fprintf(resp, "error reading request body")
	}

	cluster := &api.Cluster{}
	err = cluster.Scan(reqBody)
	if err != nil {
		log.WithError(err).Error("error parsing request body")
		fmt.Fprintf(resp, "error parsing request body")
	}
	h.logger.WithField("cluster", cluster.Name).Info("called post and parsed cluster")

	// TODO: require namespace to exist?

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err = h.db.Exec("INSERT INTO clusters (name, namespace, data) VALUES($1, $2, $3)", cluster.Name, cluster.Namespace, cluster)
	if err != nil {
		log.WithError(err).Error("error inserting into db")
		fmt.Fprintf(resp, "error inserting into db")
	}

	/*
		var newEvent event
			json.Unmarshal(reqBody, &newEvent)
			events = append(events, newEvent)
			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(newEvent)
	*/
}
