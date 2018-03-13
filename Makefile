
APP=podrick
PKG=/go/src/github.com/TimWoolford/${APP}
NAMESPACE?=monitoring

BIN=$(firstword $(subst :, ,${GOPATH}))/bin
GODEP = $(BIN)/dep
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: gobuild
gobuild: vendor
	GOOS=linux go build -v -o bin/${APP} .

.PHONY: build
build:
	docker run --rm \
	 -v "${PWD}":${PKG} \
	 -w ${PKG} \
	 golang:1.9 \
	 make gobuild

.PHONY: build-image
build-image:
	docker build -t ${APP} .

.PHONY: run
run:
	helm upgrade --install ${APP} charts/minikube --namespace ${NAMESPACE}

.PHONY: clean
clean: ; $(info $(M) cleaning…)
	@helm list -q | grep ${APP} | xargs helm delete --purge
	@docker images -q ${APP} | xargs docker rmi -f
	@rm -rf bin/*

.PHONY: vendor
vendor: .vendor

.vendor: Gopkg.toml Gopkg.lock
	command -v dep >/dev/null 2>&1 || go get github.com/golang/dep/cmd/dep
	$(GODEP) ensure -v
	@touch $@
