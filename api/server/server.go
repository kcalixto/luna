package server

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	"github.com/kcalixto/mojo-jojo/api/controllers"
	"github.com/kcalixto/mojo-jojo/api/server/router"
	"github.com/kcalixto/mojo-jojo/api/services"
)

func NewLocalServer(
	_controllers *controllers.Controller,
	_services *services.Services,
	_validator *validator.Validate,
	_config *config.Config,
) {
	_config = config.New()
	_validator = validator.New()
	_services = services.New(_validator, _config)
	_controllers = controllers.New(_services)

	ginEngine := router.NewEngine(_controllers)
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		panic("SERVER_PORT not set")
	}

	ginEngine.Run(port)
}

func NewLambdaServer(
	initialized bool,
	ginLambda *ginadapter.GinLambda,
	_controllers *controllers.Controller,
	_services *services.Services,
	_validator *validator.Validate,
	_config *config.Config,
) func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		if initialized {
			return ginLambda.Proxy(request)
		}
		// not initialized
		if _config == nil {
			_config = config.New()
		}
		if _validator == nil {
			_validator = validator.New()
		}
		if _services == nil {
			_services = services.New(_validator, _config)
		}
		if _controllers == nil {
			_controllers = controllers.New(_services)
		}

		if !initialized {
			ginEngine := router.NewEngine(_controllers)
			ginLambda = ginadapter.New(ginEngine)
			initialized = true
		}

		return ginLambda.Proxy(request)
	}
}
