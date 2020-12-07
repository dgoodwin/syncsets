# Sync Sets

Prototype for an operator like application not depending on Kubernetes CRDs or API, breaking out the SyncSets functionality from OpenShift Hive.

## Development

### Install go-swagger

Used to generate API server, go client code, and documentation.

```bash
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

Presently using "code first" with go-swagger where types/handlers have appropriate godoc annotations to generate swagger.yaml, from which virtually everything under client/ and restapi/ is generated.

### Install HTTPPie

Just an easier option than curl. Used in some commands in this README.

### Launch RabbitMQ

This project presently aims to use RabbitMQ for pub/sub consumers who wish to watch API events.

Using the [RabbitMQ Operator](https://www.rabbitmq.com/kubernetes/operator/operator-overview.html), this process involves some resources I had to patch to work on OpenShift.

```bash
kubectl apply -f manifests/rabbitmq-operator/
kubectl apply -f manifests/namespace.yaml
kubectl apply -f manifests/rabbitmq-cluster.yaml
oc adm policy add-scc-to-user rabbitmq-cluster -z rabbitmq-server
```

Once running you can check in with:

```bash
$ oc rsh rabbitmq-server-0 rabbitmqctl cluster_status
```

### Create a PostgreSQL Database

Several options here:

  1. On OpenShift: Use the OpenShift Template, create a new project in the console, select to add a database and choose postgresql.
    * Crunchy PostgreSQL operator appears much too complicated and possibly broken.
    * TODO: Try EnterpriseDB PostgreSQL operator.
  1. On plain Kubernetes/Kind: Run a postgresql pod: `kubectl create -f manifests/postgresql/postgresql.yaml`
    * This won't work on OpenShift as the official Docker images assume root.
    * TODO: Update the manifest to use the OpenShift image and have one manifest that works on both.
  1. Amazon RDS: Choose free tier and public access.

Note your password and ensure you've created a database called `syncsets`:

Establish a local port forward if running on OpenShift or Kube.

The `POSTGRES_PARAMS` env var will be used both for goose schema migrations and the api server itself.

You should be able to connect to your local database with:

```bash
kubectl port-forward svc/postgresql 5432:5432
export POSTGRES_PARAMS="user=postgres password=helloworld dbname=syncsets sslmode=disable host=localhost"
psql $POSTGRES_PARAMS
```

### Database Schema

Install goose for managing database schema migrations, and create (or update) the schema:

```bash
go get -u github.com/pressly/goose/cmd/goose
goose postgres $POSTGRES_PARAMS up
```

### Testing Locally

Ensure you have postgresql properly configured and reachable from localhost per above.

Compile your current code, generate server/client, and run the API locally:

```bash
make run
```

Push some data with httpie:

```bash
echo '{"name": "cluster1", "namespace": "foo", "kubeconfig": "foobar"}' | http POST localhost:7070/v1/clusters
```

Your database should now have an entry in the `clusters` table.

```bash
psql $POSTGRES_PARAMS -c "select * from clusters"
```

### Testing In-Cluster

WARNING: WIP, needs an update since switching to go-swagger for restapi gen.

```bash
IMG="quay.io/dgoodwin/syncsets:latest" make docker-push deploy
```

