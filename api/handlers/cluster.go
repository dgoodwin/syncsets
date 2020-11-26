package handlers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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

func (h *ClusterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("called servehttp")
	fmt.Fprintf(w, "called cluster handler")
}
