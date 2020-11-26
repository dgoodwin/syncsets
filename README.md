# Sync Sets

Prototype for an operator like application not depending on Kubernetes CRDs or API, breaking out the SyncSets functionality from OpenShift Hive.

## Development

### Create a PostgreSQL Database

Several options here:

  1. OpenShift Template: Create a new project in the console, select to add a database and choose postgresql.
    * Crunchy PostgreSQL operator appears much too complicated and possibly broken.
    * kubectl port-forward to expose your db locally.
  1. Amazon RDS: Choose free tier and public access.

Note your password and ensure you've created a database called `syncsets`:

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



