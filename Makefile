.PHONY: all get test clean build cover compile goxc bintray

VERSION=1.0.0
GO ?= go
BIN_NAME=UDP-proxy
GO_XC = ${GOPATH}/bin/goxc -os="freebsd openbsd netbsd solaris dragonfly darwin linux"
GOXC_FILE = .goxc.local.json
GITHASH=`git rev-parse HEAD`

all: clean build

get:
	${GO} get

build: get
	${GO} build -ldflags "-X main.version=${VERSION} -X main.githash=${GITHASH}" -o ${BIN_NAME} cmd/UDP-proxy/main.go;

clean:
	@rm -rf ${BIN_NAME} ${BIN_NAME}.debug *.out build debian

test: get
	${GO} test -v

cover:
	${GO} test -cover && \
	${GO} test -coverprofile=coverage.out  && \
	${GO} tool cover -html=coverage.out

compile: goxc

goxc:
	$(shell echo '{\n "ConfigVersion": "0.9",' > $(GOXC_FILE))
	$(shell echo ' "TaskSettings": {' >> $(GOXC_FILE))
	$(shell echo '  "bintray": {\n   "apikey": "$(BINTRAY_APIKEY)"' >> $(GOXC_FILE))
	$(shell echo '  }\n } \n}' >> $(GOXC_FILE))
	${GO_XC}

bintray:
	${GO_XC} bintray
