# Sync Sets

Prototype for an operator like application not depending on Kubernetes CRDs or API, breaking out the SyncSets functionality from OpenShift Hive.

## Development

All instructions assume running in an OpenShift 4 cluster.

### Launch RabbitMQ

This project presently aims to use RabbitMQ for pub/sub consumers who wish to watch API events.

Using the [RabbitMQ Operator](https://www.rabbitmq.com/kubernetes/operator/operator-overview.html), this process involves some resources I had to patch.

```bash
$ kubectl apply -f manifests/rabbitmq-operator/
$ kubectl apply -f manifests/namespace.yaml
$ kubectl apply -f manifests/rabbitmq-cluster.yaml
$ oc adm policy add-scc-to-user rabbitmq-cluster -z rabbitmq-server
```

### Create a PostgreSQL Database

Several options here:

  1. OpenShift Template: Create a new project in the console, select to add a database and choose postgresql.
    * Crunchy PostgreSQL operator appears much too complicated and possibly broken.
    * kubectl port-forward to expose your db locally.
  1. Amazon RDS: Choose free tier and public access.

Note your password and ensure you've created a database called `syncsets`:


Establish a local port forward if running on OpenShift.

```bash
$ export PGPASSWORD=MYPASS
$ psql -h localhost -U postgres -c 'create database syncsets'
```

### Database Schema

Install goose for managing database schema migrations.

```bash
$ go get -u github.com/pressly/goose/cmd/goose
```

Create a postgresql `syncsets` database on RDS or deploy one locally on kind (WARNING: does not yet work on OpenShift due to permissions):

```bash
$ kubectl create -f manifests/postgresql/postgresql.yaml
```

Apply the database schema:

```bash
$ export GOOSE_PARAMS="user=postgres dbname=syncsets sslmode=disable host=localhost password=MYPASS"
$ goose postgres $GOOSE_PARAMS up
```



### Compile the Code

```bash
$ make build
```


### Load Some Data

```bash
$ curl --header "Content-Type: application/json" --request POST -d @examples/cluster.json http://localhost:8080/clusters
$ curl --header "Content-Type: application/json" --request POST -d @examples/syncset.json http://localhost:8080/syncsets
```
