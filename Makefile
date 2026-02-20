GOLANGCI_VERSION = "v2.10.1"

all: lint test

lint: .golangci
	golangci-lint -v run

.golangci:
	which golangci-lint > /dev/null \
		&& (echo 'golangci-lint version:'; golangci-lint --version) \
		|| curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b ${GOPATH}/bin ${GOLANGCI_VERSION}

test:
	go test ./...

cover:
	go test -coverpkg=./internal/... -coverprofile=$(COVERAGE_FILE) ./... || rm -f $(COVERAGE_FILE)
	go tool cover -func $(COVERAGE_FILE)
	rm -f $(COVERAGE_FILE)
