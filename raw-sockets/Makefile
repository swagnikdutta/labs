# Makefile

LINUX_BINARY=sender-linux
DARWIN_BINARY=sender-darwin

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${LINUX_BINARY} ./cmd/sender

build-darwin:
	GOOS=darwin GOARCH=arm64 go build -o ${DARWIN_BINARY} ./cmd/sender

build:
	${MAKE} build-darwin