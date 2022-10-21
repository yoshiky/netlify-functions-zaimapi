.PHONY: build

build:
	mkdir -p functions
	go get ./...
	go install ./...
	go build -o ./functions/zaimapi ./cmd/main.go

sam:
	sam build