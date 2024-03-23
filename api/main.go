package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		ginEngine := mountRoute("hello", "get", helloWorld)
		ginLambda = ginadapter.New(ginEngine)
		initialized = true
	}

	return ginLambda.Proxy(request)
}

func helloWorld(c *gin.Context) {
	c.String(200, "Hello World!")
}

func main() {
	lambda.Start(Handler)
}

func mountRoute(path string, method string, fn gin.HandlerFunc) *gin.Engine {
	engine := buildEngine()
	group := engine.Group("/")
	setMethodHandlerForGroup(method, path, fn, group)
	return engine
}

func buildEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return engine
}

func setMethodHandlerForGroup(method string, path string, fn gin.HandlerFunc, group *gin.RouterGroup) {
	switch method {
	case "post":
		{
			group.POST(path, fn)
		}
	case "get":
		{
			group.GET(path, fn)
		}
	case "put":
		{
			group.PUT(path, fn)
		}
	case "delete":
		{
			group.DELETE(path, fn)
		}
	}
}
