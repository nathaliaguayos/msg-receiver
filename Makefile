run:
	source ./local.env && go run cmd/msg-receiver/main.go
.PHONY: run

generate:
	go generate ./...
.PHONY: run

test:
	go test -v -coverprofile=coverage.txt ./...
.PHONY: test

lint:
	go vet ./...
.PHONY: lint

htmlcoverage: test
	go tool cover -html=coverage.txt
.PHONY: htmlcoverage