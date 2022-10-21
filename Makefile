.PHONY: build

build:
	mkdir -p functions
	go get ./...
	go install ./...
	go build -o ./functions/zaimapi ./src/github.com/zaimapi/main.go

sam:
	sam build