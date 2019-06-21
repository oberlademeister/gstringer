#!/usr/bin/env bash

VERSION=`git describe --long --tags --dirty --always`
COMMIT=`git rev-parse HEAD`
BRANCH=`git rev-parse --abbrev-ref HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS="-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

if [ -z "$GOTARGET" ]; then
	GO111MODULE=on packr build -v -ldflags "${LDFLAGS}"
else
	GO111MODULE=on packr build -o ${GOTARGET} -ldflags "${LDFLAGS}"
fi

