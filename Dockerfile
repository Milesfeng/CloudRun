FROM golang

EXPOSE 8080

WORKDIR /go-hello-world

COPY go-hello-world/* .

CMD go run .