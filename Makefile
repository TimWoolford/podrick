APP=podrick
PKG=/go/src/github.com/TimWoolford/${APP}


build: vendor
	GOOS=linux go build -v -o bin/${APP} .

build2:
	docker run --rm \
	 -v "${PWD}":${PKG} \
	 -w ${PKG} \
	 golang:1.9 \
	 make build

build-image:
	docker build -t ${APP} .

run:
	helm upgrade --install ${APP} charts/minikube --namespace timtim

clean:
	helm list -q | grep ${APP} | xargs helm delete --purge
	docker images -q ${APP} | xargs docker rmi -f
	rm -rf bin/*
	rm -rf .vendor


vendor: .vendor

.vendor: Gopkg.toml Gopkg.lock
	command -v dep >/dev/null 2>&1 || go get github.com/golang/dep/cmd/dep
	dep ensure -v
	@touch $@
