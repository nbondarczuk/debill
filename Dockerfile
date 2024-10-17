FROM golang:1.21.4-alpine as builder

ENV GOPATH /go
ENV GOCACHE /go/caches/go-build
RUN apk add git alpine-sdk
ADD . /work
WORKDIR /work

# build the source
RUN make tidy
RUN make build

# use a minimal alpine image
FROM alpine:latest

# add ca-certificates in case you need them
RUN apk update && apk add ca-certificates

# set working directory
WORKDIR /work

COPY --from=builder /work/bin/debill-api-server /work/debill-api-server
COPY config/config.yaml /work/config/config.yaml
USER 1001

# run the binary
CMD ["./debill-api-server"]
