package entrypoint

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type RequestApiGateway events.APIGatewayProxyRequest
type ResponseApiGateway events.APIGatewayProxyResponse

func NewRequestByApiGateway(req RequestApiGateway) Request {
	return Request{
		Body: req.Body,
		Headers: req.Headers,
		Vars: req.PathParameters,
	}
}

func NewResponseToApiGateway(status int, response Response) ResponseApiGateway {
	b,_ := json.Marshal(response)
	return ResponseApiGateway{
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "POST,GET,OPTIONS",
			"Access-Control-Allow-Headers":     "application/json",
		},
		StatusCode: status,
		IsBase64Encoded: false,
		Body: string(b),
	}
}