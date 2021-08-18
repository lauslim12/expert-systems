.PHONY: start
start:
	go run ./cmd/expert-systems/main.go

.PHONY: build
build:
	go build -v ./cmd/expert-systems/main.go

.PHONY: format
format:
	test -z $(gofmt -l .)

.PHONY: test
test:
	go test -v -coverpkg ./internal/application ./internal/application

.PHONY: e2e
e2e:
	sh ./scripts/e2e-testing.sh
