Install the postgresql operator:

kubectl create -f https://operatorhub.io/install/postgresql.yaml

Deploy a simple postgresql server for development.

```
$ kubectl create -f manifests/postgresql/postgresql.yaml
```
