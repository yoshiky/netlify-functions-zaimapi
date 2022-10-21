package main

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	t := time.Now()
	const layout = "Now, Monday Jan 02 15:04:05 JST 2006"
	
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World! Now:" + t.Format(layout),
	}, nil
}

func main() {
	// Initiate AWS Lambda handler
	lambda.Start(Handler)
}
