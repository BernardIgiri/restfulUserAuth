FROM golang:1.9

RUN mkdir -p /go/src/application
WORKDIR /go/src/application

RUN go get github.com/tools/godep
COPY src/* /go/src/application/

RUN godep save

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN godep go build  -ldflags '-w -s' -a -installsuffix cgo -o application
