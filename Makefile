# Project Name

SHA1         		:= $(shell git rev-parse --verify --short HEAD)
MAJOR_VERSION		:= $(shell cat service.json | sed -n 's/.*"major": "\(.*\)",/\1/p')
MINOR_VERSION		:= $(shell cat service.json | sed -n 's/.*"minor": "\(.*\)"/\1/p')
INTERNAL_BUILD_ID	:= $(shell [ -z "${TRAVIS_BUILD_NUMBER}" ] && echo "local" || echo ${TRAVIS_BUILD_NUMBER})
BINARY				:= $(shell cat service.json | sed -n 's/.*"name": "\(.*\)",/\1/p')
VERSION				:= $(shell echo "${MAJOR_VERSION}_${MINOR_VERSION}_${INTERNAL_BUILD_ID}_${SHA1}")
BUILD_IMAGE			:= $(shell echo "golang:1.11.4")
PWD					:= $(shell pwd)

.PHONY: tests
tests:
	@echo "Running Tests"
	@echo "$(PWD)"

	docker run -e GO111MODULE=on --rm -t -v $(PWD):/go/src -w /go/src $(BUILD_IMAGE) go test -cover -v $(go list ./... | grep -v /example/)
	@echo "Completed tests"

.PHONY: build
build: tests
	@echo "Building in a container"

	docker run --rm -t -v $(PWD):/go/src -w /go/src $(BUILD_IMAGE) go build -x -ldflags "-X main.version=$(VERSION)" -o $(BINARY) cmd/$(BINARY)/main.go
	@echo "Executable is available at the root of cloned repo"

.PHONY: package
package: build
	@echo "Packing for Lambda"
	zip $(PROJECT).zip $(BINARY)
	rm -rf $(BINARY)

.PHONY: run
run: 
	@echo "Starting the app"
	docker run --rm -t -v $(PWD):/go/src -w /go/src $(BUILD_IMAGE) go run -ldflags "-X main.version=$(VERSION)" -v cmd/$(BINARY)/main.go --debug

