ARG GO_VERSION=1.13.15

FROM golang:${GO_VERSION}-alpine as builder
    RUN apk add --no-cache ca-certificates git tzdata
    # Configure Go
    ENV GOPATH /go
    ENV PATH /go/bin:$PATH
    ENV GOROOT /usr/local/go/
    RUN mkdir -p ${GOPATH}/src ${GOPATH}/bins
    WORKDIR $GOPATH
    # install GO
    RUN apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go

    ENV GOOS=linux
    ENV GOARCH=amd64
    RUN apk add --no-cache git

    run go get github.com/cespare/reflex
    run go get -u github.com/gin-gonic/gin
    run go get github.com/aws/aws-lambda-go/lambda

    RUN apk del .build-deps && rm -rf ~/.cache

    COPY /router /usr/local/go/src/entityValidator.com/router
    COPY /request /usr/local/go/src/entityValidator.com/request

    COPY . /go/src/entityValidator.com/

    expose 8080
    expose 80

    RUN reflex -r "\.go$$" -s -- sh -c "go run /go/src/entityValidator.com/main.go"
