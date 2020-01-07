.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./trigger/trigger
	rm -rf ./handler/handler

build:
	GOOS=linux GOARCH=amd64 go build -o handler/handler ./handler
	GOOS=linux GOARCH=amd64 go build -o trigger/trigger ./trigger

deploy: build
	sam deploy \
	--stack-name sqs-fifo \
	--capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND
