#!/bin/bash

docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd)/swagger quay.io/goswagger/swagger $*
