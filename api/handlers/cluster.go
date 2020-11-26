package handlers

import (
	//"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ClusterHandler struct {
	logger log.FieldLogger
}

func NewClusterHandler() *ClusterHandler {
	return &ClusterHandler{
		logger: log.WithField("handler", "cluster"),
	}
}

func (h *ClusterHandler) Get(resp http.ResponseWriter, req *http.Request) {
	h.logger.Info("called Get")
	fmt.Fprintf(resp, "called get cluster handler")
}

func (h *ClusterHandler) Post(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "called post cluster handler")
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.WithError(err).Error("error reading request body")
		fmt.Fprintf(resp, "error reading request body")
	}
	h.logger.WithField("body", string(reqBody)).Info("called post")

	/*
		var newEvent event
			json.Unmarshal(reqBody, &newEvent)
			events = append(events, newEvent)
			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(newEvent)
	*/
}
