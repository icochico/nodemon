# Set the GOPATH
GOPATH := ${PWD}/
export GOPATH


# This is how we want to name the binary output
BINARY=nodemon
# URLs for NATS and Go (Linux only)
NATS_URL=https://github.com/nats-io/gnatsd/releases/download/v0.9.6/gnatsd-v0.9.6-linux-amd64.zip 
GO_URL=https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz

# These are the values we want to pass for Version, BuildTime, and GitHash
VERSION=1.0.1
BUILD_TIME=`date +%FT%T%z`
GIT_HASH=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitHash=${GIT_HASH}"

.PHONY: all
all:
	go get ihmc.us/nodemon
	go build ${LDFLAGS} -o bin/${BINARY} ihmc.us/nodemon

.PHONY: go
go:
	wget ${GO_URL}
	sudo tar -C /usr/local -xzf *.tar.gz
	rm *.tar.gz

.PHONY: nats
nats:
	wget ${NATS_URL}
	mv *.zip bin/
	unzip bin/*.zip -d bin/
	rm bin/*.zip

.PHONY: clean
clean:
	go clean
	rm -rf bin/${BINARY}
