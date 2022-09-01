build:
	mkdir -p functions
	go install ./functions/zaimapi
	go build -o ./dist/zaimapi ./functinos/zaimapi