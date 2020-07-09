MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAME:=tekton-launcher

demo: install-tekton build
	$(MAKEPATH)/$(NAME) run $(MAKEPATH)/examples/example.yaml

build: fmt-and-vet test
	cd $(MAKEPATH); go build -o $(NAME) .

fmt-and-vet: 
	cd $(MAKEPATH); go fmt ./...
	cd $(MAKEPATH); go vet ./...

test:
	cd $(MAKEPATH); go test ./...

cleanup-cluster: delete-tasks-and-taskruns uninstall-tekton

install-tekton:
	-kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

uninstall-tekton:
	-kubectl delete --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

delete-tasks-and-taskruns:
	-kubectl delete taskruns.tekton.dev --all
	-kubectl delete tasks.tekton.dev --all
