package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgoodwin/syncsets/api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ClusterHandler struct {
	logger log.FieldLogger
	db     *sql.DB
}

func NewClusterHandler(db *sql.DB) *ClusterHandler {
	return &ClusterHandler{
		db:     db,
		logger: log.WithField("handler", "cluster"),
	}
}

func (h *ClusterHandler) Get(resp http.ResponseWriter, req *http.Request) {
	h.logger.Info("called Get")
	item := new(api.ClusterItem)
	err := h.db.QueryRow("SELECT id, data FROM clusters ORDER BY id DESC LIMIT 1").Scan(&item.ID, &item.Cluster)
	if err != nil {
		log.WithError(err).Error("error querying db")
	}
	jsonBytes, err := json.Marshal(item.Cluster)
	if err != nil {
		log.WithError(err).Error("error marshalling json")
	}
	fmt.Fprintf(resp, string(jsonBytes))
}

func (h *ClusterHandler) Post(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "called post cluster handler")
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

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err = h.db.Exec("INSERT INTO clusters (data) VALUES($1)", cluster)
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
