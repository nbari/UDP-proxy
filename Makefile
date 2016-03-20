.PHONY: all get test clean build cover compile goxc bintray

VERSION=0.3.0
GO ?= go
BIN_NAME=UDP-proxy
GO_XC = ${GOPATH}/bin/goxc -os="freebsd openbsd netbsd solaris dragonfly darwin linux"
GOXC_FILE = .goxc.json
GOXC_FILE_LOCAL = .goxc.local.json
GITHASH=$(shell git rev-parse HEAD)

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
	$(shell echo '{\n  "ConfigVersion": "0.9",' > $(GOXC_FILE))
	$(shell echo '  "AppName": "UDP-proxy",' >> $(GOXC_FILE))
	$(shell echo '  "ArtifactsDest": "build",' >> $(GOXC_FILE))
	$(shell echo '  "PackageVersion": "${VERSION}",' >> $(GOXC_FILE))
	$(shell echo '  "TaskSettings": {' >> $(GOXC_FILE))
	$(shell echo '    "bintray": {' >> $(GOXC_FILE))
	$(shell echo '      "downloadspage": "bintray.md",' >> $(GOXC_FILE))
	$(shell echo '      "package": "UDP-proxy",' >> $(GOXC_FILE))
	$(shell echo '      "repository": "UDP-proxy",' >> $(GOXC_FILE))
	$(shell echo '      "subject": "nbari"' >> $(GOXC_FILE))
	$(shell echo '    }\n  },' >> $(GOXC_FILE))
	$(shell echo '  "BuildSettings": {' >> $(GOXC_FILE))
	$(shell echo '    "LdFlags": "-X main.version=${VERSION} -X main.githash=${GITHASH}"' >> $(GOXC_FILE))
	$(shell echo '  }\n}' >> $(GOXC_FILE))
	$(shell echo '{\n "ConfigVersion": "0.9",' > $(GOXC_FILE_LOCAL))
	$(shell echo ' "TaskSettings": {' >> $(GOXC_FILE_LOCAL))
	$(shell echo '  "bintray": {\n   "apikey": "$(BINTRAY_APIKEY)"' >> $(GOXC_FILE_LOCAL))
	$(shell echo '  }\n } \n}' >> $(GOXC_FILE_LOCAL))
	${GO_XC}

bintray:
	${GO_XC} bintray
