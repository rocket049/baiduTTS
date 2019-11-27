#!/bin/sh
export CGO_ENABLED=0
export GOOS=windows
export GOARCH=386
LDFLAGS='-ldflags -s'
go build ${LDFLAGS}
