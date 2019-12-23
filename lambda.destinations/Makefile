.PHONY: deps clean build start-api

deps:
	go get -u ./...

clean:
	rm -rf ./hello-world/hello-world

gobuild:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
	GOOS=linux GOARCH=amd64 go build -o success/success ./success
	GOOS=linux GOARCH=amd64 go build -o failure/failure ./failure

start-api: gobuild
	sam local start-api -t ./template.yaml

package: gobuild
	sam package \
	--template-file ./template.yaml \
	--s3-bucket wb.lambda.destinations.arifact \
	--output-template-file ./packaged.yaml

deploy: package
	sam deploy \
	--template-file ./packaged.yaml \
	--stack-name lambda-destinations \
	--capabilities CAPABILITY_IAM CAPABILITY_AUTO_EXPAND

invoke-success:
	aws lambda --region eu-west-1 invoke --function-name lambda-destinations-HelloWorldFunction-CAJ9RU310Q3S --invocation-type Event --payload '{ "Success": true }' response.json

invoke-fail:
	aws lambda --region eu-west-1 invoke --function-name lambda-destinations-HelloWorldFunction-CAJ9RU310Q3S --invocation-type Event --payload '{ "Success": false }' response.json

# Local debbugging not working yet

debug-build:
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o hello-world/hello-world ./hello-world


get-dlv:
	go get -u github.com/go-delve/delve/cmd/dlv

dlv:
 	GOARCH=amd64 GOOS=linux go build -o ./dlv github.com/go-delve/delve/cmd/dlv

start-debug: debug-build
	sam local start-api -d 5986 --debugger-path . --debug-args "-delveAPI=2"