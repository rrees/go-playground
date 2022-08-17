
EXECUTABLE_NAME=dice

.PHONY: run
run: build
	./${EXECUTABLE_NAME}

.PHONY: fmt
fmt:
	go fmt

.PHONY: test
test:
	go test

build:
	go build -o ${EXECUTABLE_NAME}