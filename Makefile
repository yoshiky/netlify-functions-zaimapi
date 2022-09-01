build:
	ls -al
	cd /netlify/functions/zaimapi
	go build -o ../../../dist/zaimapi ../../functions/zaimapi