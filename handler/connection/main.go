package main

import (
	"context"
	"lambda/gateway"
	"lambda/usecase"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	connectionRepository usecase.IConnectionRepository
	messageSender        usecase.IMessageSender
	usecases             usecase.IUsecase
)

type Response = events.APIGatewayProxyResponse

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	connectionRepository = gateway.NewConnectionRepository(
		dynamodb.NewFromConfig(cfg),
		"connectionId",
		os.Getenv("CONNECTION_TABLE_NAME"),
	)

	usecases = usecase.NewUsecase(connectionRepository, messageSender)

	lambda.Start(handler)
}

func handler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (Response, error) {
	log.Println("requestId", req.RequestContext.RequestID)
	log.Println("connectionId", req.RequestContext.ConnectionID)

	if err := usecases.OnConnect(req.RequestContext.ConnectionID); err != nil {
		log.Println(err)
		// 雑にすべて500エラーを返します
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	return Response{StatusCode: http.StatusOK}, nil
}
