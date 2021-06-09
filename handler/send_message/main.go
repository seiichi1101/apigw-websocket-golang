package main

import (
	"context"
	"encoding/json"
	"lambda/gateway"
	"lambda/usecase"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	endpoint             url.URL
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

	endpoint.Scheme = "https"
	endpoint.Path = "dev"
	endpoint.Host = os.Getenv("APIGW_HOST")
	endpointResolver := apigatewaymanagementapi.EndpointResolverFromURL(endpoint.String())

	messageSender = gateway.NewMessageSender(
		apigatewaymanagementapi.NewFromConfig(
			cfg,
			apigatewaymanagementapi.WithEndpointResolver(endpointResolver),
		),
	)

	usecases = usecase.NewUsecase(connectionRepository, messageSender)

	lambda.Start(handler)
}

type body struct {
	Message  string `json:"message"`
	TargetId string `json:"targetId"`
}

func handler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (Response, error) {
	log.Println("requestId", req.RequestContext.RequestID)
	log.Println("connectionId", req.RequestContext.ConnectionID)

	body := body{}
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		log.Println(err)
		// 雑にすべて500エラーを返します
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	log.Println("body", body)
	if err := usecases.SendMessage(body.TargetId, body.Message); err != nil {
		log.Println(err)
		// 雑にすべて500エラーを返します
		return Response{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	return Response{StatusCode: http.StatusOK}, nil
}
