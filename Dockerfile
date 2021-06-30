FROM golang:latest
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .
RUN apt-get -y update && apt-get -y install git
RUN go get "github.com/gorilla/mux"
ENTRYPOINT go run main.go
