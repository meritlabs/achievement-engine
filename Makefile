GO_PATH=$(GOPATH)
GO_PATH?=/tmp/go
GO_SRC=$(GO_PATH)/src
PACKAGE_PATH=github.com/meritlabs
PACKAGE=$(PACKAGE_PATH)/achievement-engine
SRC=$(GO_SRC)/$(PACKAGE)

GO=go
GO_FMT=$(GO) fmt

AE_BIN=achievement-engine
AE_MIGRATIONS_BIN=achievement-engine-migrations

.PHONY: build
build: build-achievement-engine build-achievement-engine-migrations

.PHONY: build-achievement-engine
build-achievement-engine:
	$(GO) build -o $(AE_BIN) cmd/api/main.go

.PHONY: build-achievement-engine-migrations
build-achievement-engine-migrations:
	$(GO) build -o $(AE_MIGRATIONS_BIN) cmd/migrations/main.go

.PHONY: clean
clean:
	rm -r vendor $(AE_BIN) $(AE_MIGRATIONS_BIN)

.PHONY: bootstrap
bootstrap:
	$(GO) mod vendor
