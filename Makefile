build:
	go build

.PHONY: dependencies
dependencies:
	go get -v -t -d ./...

.PHONY: build_aws_lambda
build_aws_lambda: dependencies_aws_lambda
	env GOOS=linux GOARCH=amd64 go build -tags aws_lambda

.PHONY: dependencies_aws_lambda
dependencies_aws_lambda:
	go get -v -t -d -tags aws_lambda ./...
