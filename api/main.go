package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/luna/api/config"
	controllers "github.com/kcalixto/luna/api/controllers"
	"github.com/kcalixto/luna/api/data/repository"
	"github.com/kcalixto/luna/api/server"
	"github.com/kcalixto/luna/api/services"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

// Lambdas can keep these values "in memory", so we don't need to re-create them every time a new lambda is called!
var inMemory_config *config.Config
var inMemory_validator *validator.Validate
var inMemory_repository *repository.RepositoryManager
var inMemory_services *services.Services
var inMemory_controllers *controllers.Controller

func main() {
	if os.Getenv("ENV") == "local" {
		server.NewLocalServer(
			inMemory_controllers,
			inMemory_services,
			inMemory_validator,
			inMemory_config,
			inMemory_repository,
		)
	} else {
		lambda.Start(
			server.NewLambdaServer(
				initialized,
				ginLambda,
				inMemory_controllers,
				inMemory_services,
				inMemory_validator,
				inMemory_config,
				inMemory_repository,
			),
		)
	}
}
