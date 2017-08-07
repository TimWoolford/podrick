APP=podrick
PKG=/go/src/github.com/TimWoolford/${APP}


build:
	go get ./cmd/${APP}
	GOOS=linux go build -o bin/${APP} ./cmd/${APP}

build2:
	docker run --rm \
	 -v "${PWD}":${PKG} \
	 -w ${PKG} golang:1.8 \
	 go get ./cmd/${APP} && go build -o bin/${APP} \
	 ./cmd/${APP}

build-image:
	docker build -t ${APP} .

run:
	helm upgrade --install ${APP} charts/minikube

clean:
	helm list -q | grep ${APP} | xargs helm delete --purge
	docker images -q ${APP} | xargs docker rmi
	rm -rf bin/*
