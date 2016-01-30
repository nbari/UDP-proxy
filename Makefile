.PHONY: all get test clean build cover compile goxc bintray

GO ?= go
BIN_NAME=UDP-proxy
GO_XC = ${GOPATH}/bin/goxc -os="freebsd openbsd netbsd solaris dragonfly darwin linux"
GOXC_FILE = .goxc.local.json

all: clean build

get:
	${GO} get

build: get
# make build DEBUG=true
	@if test -n "${DEBUG}"; then \
	${GO} get -u github.com/mailgun/godebug; \
	${GOPATH}/bin/godebug build -instrument="github.com/nbari/UDP-proxy/..." -o ${BIN_NAME}.debug cmd/UDP-proxy/main.go; \
	else \
	${GO} build -o ${BIN_NAME} cmd/UDP-proxy/main.go; \
	fi;

clean:
	@rm -f ${BIN_NAME} ${BIN_NAME}.debug *.out

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
