
IMG ?= quay.io/dgoodwin/syncsets:latest

.PHONY: build
build: swagger-gen
	go build -o bin/syncsets-server github.com/dgoodwin/syncsets/cmd/syncsets-server
	#go build -o bin/syncsets-controllers github.com/dgoodwin/syncsets/cmd/syncsets-controllers

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

.PHONY: swagger-spec
swagger-spec:
	swagger generate spec -o ./swagger.yaml -m -q

.PHONY: swagger-validate
swagger-validate:
	swagger validate swagger.yaml -q

.PHONY: swagger-gen
swagger-gen: swagger-spec
	rm -rf restapi/operations
	swagger generate server -A syncsets -f ./swagger.yaml -q
	#swagger generate client -A syncsets -f ./swagger.yaml

.PHONY: install
install: swagger-gen
	go install ./cmd/syncsets-server

.PHONY: run
run: install
	syncsets-server --port=7070


