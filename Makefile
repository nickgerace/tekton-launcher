MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAME:=tekton-launcher

build: cleanup
	cd $(MAKEPATH); go build -o $(NAME) .

cleanup: 
	cd $(MAKEPATH); go fmt ./...
	cd $(MAKEPATH); go vet ./...

install-pipelines:
	kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

delete-pipelines:
	-kubectl delete --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

create-namespace:
	kubectl apply -f $(MAKEPATH)/namespace.yaml

delete-namespace:
	-kubectl delete -f $(MAKEPATH)/namespace.yaml

install-example:
	kubectl apply -f $(MAKEPATH)/example.yaml

delete-example:
	-kubectl delete taskruns.tekton.dev --all
	-kubectl delete tasks.tekton.dev --all
