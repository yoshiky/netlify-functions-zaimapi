package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoshiky/zaimapi/cmd/driver"
)

func main() {
	lambda.Start(driver.Handler)
}
