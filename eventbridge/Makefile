.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./orders/orders

build:
	GOOS=linux GOARCH=amd64 go build -o orders/orders ./orders
	GOOS=linux GOARCH=amd64 go build -o forecast/forecast ./forecast

deploy: build
	sam deploy \
	--stack-name eventbridge-sample \
	--capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND