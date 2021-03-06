# Project Name

SHA1         		:= $(shell git rev-parse --verify --short HEAD)
MAJOR_VERSION		:= $(shell cat service.json | sed -n 's/.*"major": "\(.*\)",/\1/p')
MINOR_VERSION		:= $(shell cat service.json | sed -n 's/.*"minor": "\(.*\)"/\1/p')
INTERNAL_BUILD_ID	:= $(shell [ -z "${TRAVIS_BUILD_NUMBER}" ] && echo "local" || echo ${TRAVIS_BUILD_NUMBER})
BINARY				:= $(shell cat service.json | sed -n 's/.*"name": "\(.*\)",/\1/p')
VERSION				:= $(shell echo "${MAJOR_VERSION}.${MINOR_VERSION}.${INTERNAL_BUILD_ID}-${SHA1}")
BUILD_IMAGE			:= $(shell echo "golang:1.11.4")
PWD					:= $(shell pwd)

.DEFAULT_GOAL := package

.PHONY: version
version:
	@echo "Setting build to Version: v$(VERSION)" 
	$(shell echo v$(VERSION) > VERSION.txt)

.PHONY: test
test: version
	@echo "Running tests in a container"
	docker run -e GO111MODULE=on --rm -t -v $(PWD):/usr/src/myapp -w /usr/src/myapp $(BUILD_IMAGE) sh -c "go test -cover -v ./... -count=1"
	@echo "Completed tests"

.PHONY: build
build: test
	@echo "Building in a container"

	docker run -e GO111MODULE=on --rm -t -v $(PWD):/usr/src/myapp -w /usr/src/myapp $(BUILD_IMAGE) sh -c "go build -x -ldflags '-X main.version=$(VERSION)' -o $(BINARY) cmd/$(BINARY)/main.go"
	@echo "Executable is available at the root of cloned repo"

.PHONY: package
package: build
	@echo "Packing binary in zip file for Lambda deployment"
	zip $(BINARY).zip $(BINARY)
	rm -rf $(BINARY)

# Local testing
.PHONY: run
run: 
	@echo "Starting the app"
	docker run --rm -t -v $(PWD):/usr/src/myapp -w /usr/src/myapp $(BUILD_IMAGE) printenv; go run -ldflags "-X main.version=$(VERSION)" -v cmd/$(BINARY)/main.go --debug
