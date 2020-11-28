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

type SyncSetHandler struct {
	logger log.FieldLogger
	db     *sql.DB
}

func NewSyncSetHandler(db *sql.DB) *SyncSetHandler {
	return &SyncSetHandler{
		db:     db,
		logger: log.WithField("handler", "syncset"),
	}
}

func (h *SyncSetHandler) Get(resp http.ResponseWriter, req *http.Request) {
	h.logger.Info("called Get")
	item := new(api.SyncSetItem)
	err := h.db.QueryRow("SELECT id, data FROM syncsets ORDER BY id DESC LIMIT 1").Scan(&item.ID, &item.SyncSet)
	if err != nil {
		log.WithError(err).Error("error querying db")
	}
	jsonBytes, err := json.Marshal(item.SyncSet)
	if err != nil {
		log.WithError(err).Error("error marshalling json")
	}
	fmt.Fprintf(resp, string(jsonBytes))
}

func (h *SyncSetHandler) Post(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "called post cluster handler")
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithError(err).Error("error reading request body")
		fmt.Fprintf(resp, "error reading request body")
	}

	ss := &api.SelectorSyncSet{}
	err = ss.Scan(reqBody)
	if err != nil {
		log.WithError(err).Error("error parsing request body")
		fmt.Fprintf(resp, "error parsing request body")
	}
	h.logger.WithField("syncset", ss.Name).Info("called post and parsed syncset")

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err = h.db.Exec("INSERT INTO syncsets (name, namespace, data) VALUES($1, $2, $3)", ss.Name, ss.Namespace, ss)
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
