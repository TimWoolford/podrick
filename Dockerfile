FROM golang:1.8

WORKDIR /app
ADD bin/podrick /app

CMD ["./podrick"]