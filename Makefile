# Project Name

SHA1         		:= $(shell git rev-parse --verify --short HEAD)
MAJOR_VERSION		:= $(shell cat service.json | sed -n 's/.*"major": "\(.*\)",/\1/p')
MINOR_VERSION		:= $(shell cat service.json | sed -n 's/.*"minor": "\(.*\)"/\1/p')
INTERNAL_BUILD_ID	:= $(shell [ -z "${BUILD_ID}" ] && echo "local" || echo ${BUILD_ID})
BINARY				:= $(shell cat service.json | sed -n 's/.*"name": "\(.*\)",/\1/p')
VERSION				:= $(shell echo "${MAJOR_VERSION}_${MINOR_VERSION}_${INTERNAL_BUILD_ID}_${SHA1}")
PWD					:= $(shell pwd)
CONFIG_TEMPLATE     := $(shell echo "stack/config.toml")

DOT:= .
DASH:= -
# replace . with -
PROJECT				:= $(subst $(DOT),$(DASH),$(BINARY))
BUILD_IMAGE			:= $(shell echo "${PROJECT}-build")

.PHONY: package
package: build
	@echo "Packing for Lambda"
	zip $(PROJECT).zip $(BINARY)
	rm $(BINARY)

.PHONY: build
build: tests
	@echo "Building in a container"

	docker run --rm --name=$(BUILD_IMAGE) -t -v $(PWD):/go/src -w /go/src $(BUILD_IMAGE) go build -x -ldflags "-X main.version=$(VERSION)" -o $(BINARY) cmd/$(BINARY)/main.go
	@echo "Executable is available at the root of cloned repo"

.PHONY: tests
tests: setup
	@echo "Running Tests"

	docker run --rm --name=$(BUILD_IMAGE) -t -v $(PWD):/go/src -w /go/src $(BUILD_IMAGE) go test -cover -v ./...
	@echo "Completed tests"


.PHONY: setup
setup:
	@echo "Setting up environment"
	@echo "$(VERSION)"

	@echo $(VERSION)
	@echo $(BINARY)

	docker rmi -f $(BUILD_IMAGE)
	docker build -t=$(BUILD_IMAGE) .

.PHONY: run
run: setup 
	@echo "Starting the app"
	docker run --rm --name=$(BUILD_IMAGE) -t -v $(PWD):/go/src -w /go/src $(BUILD_IMAGE) go run -ldflags "-X main.version=$(VERSION)" -v cmd/$(BINARY)/main.go --debug

