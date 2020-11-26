# Sync Sets

Prototype for an operator like application not depending on Kubernetes CRDs or API, breaking out the SyncSets functionality from OpenShift Hive.

## Development

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
$ export GOOSE_PARAMS="user=postgres dbname=syncsets sslmode=disable host=syncsets-1.akjsdhayusdh.us-east-2.rds.amazonaws.com password=databasepassword"
$ goose postgres $GOOSE_PARAMS up
```



