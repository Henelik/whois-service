package main

import (
	"encoding/json"
	"log"

	"Henelik/whois-service/whois"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler function for AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext)

	var reqBody map[string]string

	// check JSON is valid
	if json.Valid([]byte(request.Body)) {
		err := json.Unmarshal([]byte(request.Body), &reqBody)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       err.Error(),
				StatusCode: 500,
			}, err
		}
	} else { // unmarshal valid JSON
		return events.APIGatewayProxyResponse{
			Body:       "Invalid json format",
			StatusCode: 400,
		}, nil
	}

	// ensure a specified domain field in request body
	domain, ok := reqBody["domain"]
	if !ok {
		return events.APIGatewayProxyResponse{
			Body:       "Please specify a domain in the body of your request",
			StatusCode: 400,
		}, nil
	}

	// look up whois info on domain
	respBody, err := whois.Whois(domain, "10s")
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, err
	}

	// successful return
	return events.APIGatewayProxyResponse{
		Body:       respBody,
		StatusCode: 200,
	}, nil
}

func main() {
	log.Printf("Start lambda")
	lambda.Start(Handler)
}
