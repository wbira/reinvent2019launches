.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./normal/normal
	rm -rf ./provisioned/provisioned

build:
	GOOS=linux GOARCH=amd64 go build -o normal/normal ./normal
	GOOS=linux GOARCH=amd64 go build -o provisioned/provisioned ./provisioned

deploy: build
	sam deploy \
	--stack-name lambda-provisioned-concurrency \
	--guided \
	--capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND
