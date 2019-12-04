CIRCLE_BUILD_NUM ?= dev

APP=podrick
PKG=/go/src/github.com/TimWoolford/${APP}
TAG=timwoolford/${APP}:0.1.$(CIRCLE_BUILD_NUM)

BIN=$(firstword $(subst :, ,${GOPATH}))/bin
GODEP = $(BIN)/dep
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: gobuild
gobuild: vendor ; $(info $(M) building…)
	$Q GOOS=linux go build -v -o bin/${APP} ./cmd/podrick

.PHONY: gotest
gotest: gobuild ; $(info $(M) running tests…)
	$Q go test ./...

.PHONY: build
build:
	docker run --rm \
	 -v "${PWD}":${PKG} \
	 -w ${PKG} \
	 golang:1.12 \
	 make gobuild

.PHONY: build-image
build-image:
	docker build -t ${TAG} .

.PHONY: push-image
push-image:
	docker push ${TAG}

.PHONY: clean-vendor
clean-vendor:
	rm -rf .vendor vendor Gopkg.lock

.PHONY: clean
clean: ; $(info $(M) cleaning…)
	$Q docker images -q ${TAG} | xargs docker rmi -f
	$Q rm -rf bin/*


vendor: Gopkg.toml Gopkg.lock ; $(info $(M) retrieving dependencies…)
	$Q command -v $(GODEP) >/dev/null 2>&1 || go get github.com/golang/dep/cmd/dep
	$Q $(GODEP) ensure -v
	$Q touch $@


clean-minikube:
	helm delete ${APP} --purge

.PHONY: deploy-minikube
deploy-minikube:
	helm upgrade --tiller-namespace kube-system --install ${APP} charts/minikube --namespace monitoring