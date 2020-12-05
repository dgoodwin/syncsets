
IMG ?= quay.io/dgoodwin/syncsets:latest

.PHONY: build
build:
	go build -o bin/syncsets-api github.com/dgoodwin/syncsets/cmd/syncsets-api
	go build -o bin/syncsets-controllers github.com/dgoodwin/syncsets/cmd/syncsets-controllers

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

# Overwrites swagger.yaml. Use to go back to code first, but for now I am attempting to
# stick with spec first.
#.PHONY: swagger-spec
#swagger-spec:
#	swagger generate spec -o ./swagger.yaml -m

.PHONY: swagger-validate
swagger-validate:
	swagger validate swagger.yaml

.PHONY: swagger-gen
swagger-gen:
	rm -rf restapi/operations
	swagger generate server -A syncsets -f ./swagger.yaml
	swagger generate client -A syncsets -f ./swagger.yaml

.PHONY: install
install: swagger-gen
	go install ./cmd/syncsets-server

.PHONY: run
run: install
	syncsets-server --port=7070


