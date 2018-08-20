package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type SlackMessage struct {
}

func buildMessage(description *cloudformation.DescribeStacksOutput) (SlackMessage, error) {
	var slackMessage SlackMessage
	return slackMessage, nil
}

func HandleRequest() (events.APIGatewayProxyResponse, error) {
	client := cloudformation.New(
		session.New(aws.NewConfig().WithRegion("ap-northeast")), nil,
	)

	description, err := client.DescribeStacks(nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500, Body: "Failed to describe stacks",
		}, err
	}

	slackMessage, err := buildMessage(description)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500, Body: "Failed to build slack message",
		}, err
	}

	body, err := json.Marshal(slackMessage)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500, Body: "Failed to marshal slack message",
		}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
