.PHONY: lint

lint: ensure-linter-exists
	golangci-lint run

LINTER_VERSION=v1.18.0
ensure-linter-exists:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin $(LINTER_VERSION)
