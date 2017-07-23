APP=podrick
PKG=/go/src/github.com/TimWoolford/${APP}


build:
	docker run --rm \
	 -v "${PWD}":${PKG} \
	 -w ${PKG} golang:1.8 \
	 go build -o bin/${APP} \
	 ./cmd/${APP}


build-image:
	docker build -t ${APP} .

run:
	docker run -it --rm  -p 8082:8082 --name ${APP} ${APP}