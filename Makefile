GO_PATH=$(GOPATH)
GO_PATH?=/tmp/go
GO_SRC=$(GO_PATH)/src
PACKAGE_PATH=github.com/meritlabs
PACKAGE=$(PACKAGE_PATH)/achievement-engine
SRC=$(GO_SRC)/$(PACKAGE)

GO=go
GO_FMT=$(GO) fmt

.PHONY: build
build: build-achievement-engine build-achievement-engine-migrations

.PHONY: build-achievement-engine
build-achievement-engine:
	cd "$(SRC)" && $(GO) build -o achievement-engine api/main.go

.PHONY: build-achievement-engine-migrations
build-achievement-engine-migrations:
	cd "$(SRC)" && $(GO) build -o achievement-engine-migrations migrations/main.go

.PHONY: clean
clean:
	rm -rf  vendor
	rm -rf $(SRC)

.PHONY: bootstrap
bootstrap:
	if [ ! -d "$(SRC)" ]; then mkdir -p "$(GO_SRC)/$(PACKAGE_PATH)" && ln -s "$(PWD)" "$(SRC)" ; fi
	cd "$(SRC)" && dep ensure
