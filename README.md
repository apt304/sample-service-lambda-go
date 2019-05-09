# Sample API Service Lambda
Stop writing code that only runs on AWS Lambda!

This is a sample repo that demonstrates ideas in this [presentation](presentation.pdf).

This project uses go build tags to differentiate what main function to leverage. We have main_standalone.go and main_aws_lambda.go. When running make build main_standalone.go is used. When running make build_aws_lambda main_aws_lambda.go is used.

## Standalone Build
Build standalone go binary
```
make build
```

Run it
```
./sample-service-lambda-go
```

Test it
```
curl http://localhost:8080/hello
```

## AWS Lambda Build
Build AWS Lambda targeted go binary
```
make build_aws_lambda
```

Deploy to AWS Lambda (uses serverless but is not required)
```
serverless deploy
```

Test it
```
curl ENDPOINT_RETURNED_FROM_SERVERLESS/hello
```
