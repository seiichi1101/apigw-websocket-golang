package gateway

import (
	"context"
	"lambda/domain/connection"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

func NewMessageSender(client *apigatewaymanagementapi.Client) *MessageSender {
	return &MessageSender{client}
}

type MessageSender struct {
	client *apigatewaymanagementapi.Client
}

func (m *MessageSender) SendMessage(id connection.Id, message string) error {
	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(id.String()),
		Data:         []byte(message),
	}
	_, err := m.client.PostToConnection(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}
