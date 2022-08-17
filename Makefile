
.PHONY: run
run:
	go run *.go

.PHONY: fmt
fmt:
	go fmt

.PHONY: test
test:
	go test