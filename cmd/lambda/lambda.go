package main

import (
	"context"
	_ "embed"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginAdapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"go_api_boilerplate/infra/router"
)

var ginLambda *ginAdapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func exec() {

	// init router
	r := router.NewRouter()

	ginLambda = ginAdapter.New(r.API)
	lambda.Start(Handler)
}
