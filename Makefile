
IMG ?= quay.io/dgoodwin/syncsets:latest

.PHONY: build
build:
	go build -o bin/syncsets-api github.com/dgoodwin/syncsets/api/cmd
	go build -o bin/syncsets-controllers github.com/dgoodwin/syncsets/controllers/cmd

.PHONY: docker-push
docker-push: build
	docker build -t ${IMG} .
	docker push ${IMG}

.PHONY: deploy
deploy:
	kubectl apply -f manifests/
	oc patch deployment syncsets-api --type='json' -p='[ { op: "replace", path: "/spec/template/spec/containers/0/image", value: "${IMG}" }  ]'
	oc delete pod -l app=syncsets-api --wait=false
	oc delete pod -l app=syncsets-controllers --wait=false

