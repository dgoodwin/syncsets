package handlers

import (
	"database/sql"
	//"encoding/json"
	"github.com/dgoodwin/syncsets/api"
	"github.com/dgoodwin/syncsets/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// OK indicates that the HTTP request was successful.
//
// swagger:response
type OK struct {
	ResponseCode int
}

// BadRequest indicates that there was an error in
// the HTTP request.
//
// swagger:response
type BadRequest struct {
	ResponseCode int
}

// swagger:route POST /clusters createCluster
//
// Create a Cluster resource.
//
// Some test description that should be expanded on someday.
//
// Schemes: http
//
// Responses:
//   200: OK
//   400: BadRequest
type CreateClusterHandler struct {
	logger   log.FieldLogger
	db       *sql.DB
	registry *api.Registry
}

func NewCreateClusterHandler(db *sql.DB) *CreateClusterHandler {
	return &CreateClusterHandler{
		db:       db,
		logger:   log.WithField("handler", "cluster"),
		registry: api.NewRegistry(),
	}
}

func GetResourceType(req *http.Request) (string, error) {
	return req.URL.Path[1:], nil
}

/*
func (h *CreateClusterHandler) Handle(params clusters.GetClustersParams) middleware.Responder {
	h.logger.Info("called get handler")
	resource := "clusters"
	h.logger.Infof("working with resource: %s", resource)
	row := h.db.QueryRow(
		fmt.Sprintf("SELECT data FROM %s ORDER BY id DESC LIMIT 1", resource))

	c := &models.Cluster{}
	err := row.Scan(c)
	if err != nil {
		log.WithError(err).Error("error querying db")
		return clusters.NewGetClustersDefault(500)
	}
	return clusters.NewGetClustersOK().WithPayload([]*models.Cluster{c})
}
*/

func (h *CreateClusterHandler) Handle(params operations.CreateClusterParams) middleware.Responder {
	log.Info("called CreateClusterHandler")

	reqBody, err := ioutil.ReadAll(params.HTTPRequest.Body)
	if err != nil {
		log.WithError(err).Error("error reading request body")
		return operations.NewCreateClusterBadRequest().WithResponseCode(500)
	}

	cluster := &api.Cluster{}
	err = cluster.Scan(reqBody)
	if err != nil {
		log.WithError(err).Error("error parsing request body")
		return operations.NewCreateClusterBadRequest().WithResponseCode(500)
	}
	h.logger.WithField("cluster", cluster.Name).Info("called post and parsed cluster")

	// TODO: require namespace to exist?

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err = h.db.Exec("INSERT INTO clusters (name, namespace, data) VALUES($1, $2, $3)", cluster.Name, cluster.Namespace, cluster)
	if err != nil {
		log.WithError(err).Error("error inserting into db")
		return operations.NewCreateClusterBadRequest().WithResponseCode(500)
	}
	return operations.NewCreateClusterOK()
}
