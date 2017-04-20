#!/usr/bin/env bash
set -e

VERSION=$(grep -i PackageVersion .goxc.json | sed 's/[^[:digit:]\.]//g')
TAG="benton/htest:${VERSION}"
#goxc

echo "Building static linux binary for htest, v${VERSION}..."
mkdir -p pkg
buildcmd='CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o pkg/htest'
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" \
  -w /app golang:1.8 sh -c "$buildcmd"

echo "Building docker image for htest..."
docker build --no-cache=true --tag $TAG .
rm -f pkg/htest

echo "Done building ${TAG}."
exit 0
