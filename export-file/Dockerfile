FROM alpine:latest@sha256:0a4eaa0eecf5f8c050e5bba433f58c052be7587ee8af3e8b3910ef9ab5fbe9f5

RUN apk add --no-cach bash

ENV GO_VERSION=1.22.3

RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

ENV GOROOT=/usr/local/go

ENV GOPATH=/go

ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

WORKDIR /./Dockerize

COPY . .

ENTRYPOINT [ "go", "run", "." ]