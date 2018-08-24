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
	Text         string       `json:"text"`
	ResponseType string       `json:"response_type"`
	Attachments  []Attachment `json:"attachments"`
}

type Attachment struct {
	Title string `json:"title"`
}

func buildMessage(description *cloudformation.DescribeStacksOutput) (SlackMessage, error) {
	attachments := make([]Attachment, len(description.Stacks))
	nestedStackCount := 0
	for _, stack := range description.Stacks {
		// Nested stacks are not added to message
		if stack.ParentId != nil || stack.RootId != nil {
			nestedStackCount++
			continue
		}

		attachments = append(attachments, Attachment{
			Title: *stack.StackName,
		})
	}

	message := SlackMessage{
		Text:         "List of stacks",
		ResponseType: "ephemeral",
		Attachments:  attachments[nestedStackCount:], // Omit empty elements
	}

	return message, nil
}

func HandleRequest() (events.APIGatewayProxyResponse, error) {
	client := cloudformation.New(
		session.New(aws.NewConfig().WithRegion("ap-northeast-1")), nil,
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
