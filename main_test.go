package main

import (
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandlerPositive(t *testing.T) {
	resp, err := Handler(events.APIGatewayProxyRequest{
		RequestContext: events.APIGatewayProxyRequestContext{},
		Body:           `{"domain": "fishtech.group"}`,
	})

	assert.Nil(t, err)
	assert.IsType(t, events.APIGatewayProxyResponse{}, resp)
	assert.Equal(t, 200, resp.StatusCode)

	expectedDomain := "Domain Name: fishtech.group"
	assert.Equal(t, expectedDomain, strings.Split(resp.Body, "\r\n")[0])
}

func TestHandlerInvalidJSON(t *testing.T) {
	resp, err := Handler(events.APIGatewayProxyRequest{
		RequestContext: events.APIGatewayProxyRequestContext{},
		Body:           `{"domain": "fishtech.group"}]`,
	})

	assert.Nil(t, err)
	assert.IsType(t, events.APIGatewayProxyResponse{}, resp)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Invalid json format", resp.Body)
}

func TestHandlerNoDomain(t *testing.T) {
	resp, err := Handler(events.APIGatewayProxyRequest{
		RequestContext: events.APIGatewayProxyRequestContext{},
		Body:           `{"url": "fishtech.group"}`,
	})

	assert.Nil(t, err)
	assert.IsType(t, events.APIGatewayProxyResponse{}, resp)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Please specify a domain in the body of your request", resp.Body)
}

// More tests *could* be written for all whois package errors.
// This is sufficient to ensure that whois errors are properly passed through the Lambda.
func TestHandlerWhoisTLDError(t *testing.T) {
	resp, err := Handler(events.APIGatewayProxyRequest{
		RequestContext: events.APIGatewayProxyRequestContext{},
		Body:           `{"domain": "fishtech.grou"}`,
	})

	assert.Equal(t, "No whois server found for domain fishtech.grou", err.Error())
	assert.IsType(t, events.APIGatewayProxyResponse{}, resp)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "No whois server found for domain fishtech.grou", resp.Body)
}
