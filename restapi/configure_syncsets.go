// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/dgoodwin/syncsets/handlers"
	"github.com/dgoodwin/syncsets/restapi/operations"
	"github.com/dgoodwin/syncsets/restapi/operations/clusters"
)

//go:generate swagger generate server --target ../../syncsets --name Syncsets --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.SyncsetsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SyncsetsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	log.SetLevel(log.InfoLevel)
	log.Info("running syncsets-api")

	db, err := sql.Open("postgres", "user=postgres password=helloworld host=localhost dbname=syncsets sslmode=disable")
	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}
	log.Info("database connection established")

	/*
		conn, err := amqp.Dial("amqp://ffY5PQ_tMehsn2tryfCDvuEvVDIBoLYu:Sp7vDEG8J_E62XFi-6r3XWBQJJi0T1Sy@rabbitmq:5672/")
		if err != nil {
			log.WithError(err).Fatal("error connecting to to rabbitmq")
		}
		defer conn.Close()
		log.Info("rabbitmq connection established")
		ch, err := conn.Channel()
		if err != nil {
			log.WithError(err).Fatal("error establishing rabbitmq channel")
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"clusters", // name
			false,      // durable
			false,      // delete when unused
			false,      // exclusive
			false,      // no-wait
			nil,        // arguments
		)
		if err != nil {
			log.WithError(err).Fatal("error declaring rabbitmq queue")
		}
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("hello world"),
			})
		if err != nil {
			log.WithError(err).Error("error publishing test message")
		}
	*/

	clusterHandler := handlers.NewClusterHandler(db)

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	if api.ClustersGetClustersHandler == nil {
		api.ClustersGetClustersHandler = clusterHandler
	}
	if api.ClustersDeleteHandler == nil {
		api.ClustersDeleteHandler = clusters.DeleteHandlerFunc(func(params clusters.DeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation clusters.Delete has not yet been implemented")
		})
	}
	if api.ClustersUpdateHandler == nil {
		api.ClustersUpdateHandler = clusters.UpdateHandlerFunc(func(params clusters.UpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation clusters.Update has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
