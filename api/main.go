package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	controllers "github.com/kcalixto/mojo-jojo/api/controllers"
	"github.com/kcalixto/mojo-jojo/api/server"
	"github.com/kcalixto/mojo-jojo/api/services"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

var _config *config.Config
var _validator *validator.Validate
var _services *services.Services
var _controllers *controllers.Controller

func main() {
	if os.Getenv("ENV") == "local" {
		server.NewLocalServer(
			_controllers,
			_services,
			_validator,
			_config,
		)
	} else {
		lambda.Start(
			server.NewLambdaServer(
				initialized,
				ginLambda,
				_controllers,
				_services,
				_validator,
				_config,
			),
		)
	}
}
