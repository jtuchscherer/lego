FROM alpine:3.6

ENV GOPATH /go
ENV LEGO_VERSION acmev2

RUN apk update && apk add --no-cache --virtual run-dependencies ca-certificates && \
    apk add --no-cache --virtual build-dependencies go git musl-dev && \
    go get -u github.com/jtuchscherer/lego && \
    cd ${GOPATH}/src/github.com/jtuchscherer/lego && \
    git checkout ${LEGO_VERSION} && \
    go build -o /usr/bin/lego . && \
    apk del build-dependencies && \
    rm -rf ${GOPATH}

ENTRYPOINT [ "/usr/bin/lego" ]
