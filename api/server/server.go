package server

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kcalixto/mojo-jojo/api/config"
	"github.com/kcalixto/mojo-jojo/api/controllers"
	"github.com/kcalixto/mojo-jojo/api/data/repository"
	"github.com/kcalixto/mojo-jojo/api/server/router"
	"github.com/kcalixto/mojo-jojo/api/services"
)

func endIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewLocalServer(
	controllersPointer *controllers.Controller,
	servicesPointer *services.Services,
	validatorPointer *validator.Validate,
	configPointer *config.Config,
	repoPointer *repository.RepositoryManager,
) {
	var err error
	lambdaContext := context.Background()
	// In local development all these pointers are nil, but go linter will say that they are being replaced before it's first use
	// So we need to check if they are nil before creating a new instance to avoid this warn message!
	if configPointer == nil {
		configPointer = config.New()
		endIfError(err)
	}
	if validatorPointer == nil {
		validatorPointer = validator.New()
		endIfError(err)
	}
	if repoPointer == nil {
		repoPointer = repository.New(context.Background(), configPointer, validatorPointer)
		endIfError(err)
	}
	if servicesPointer == nil {
		servicesPointer, err = services.New(validatorPointer, configPointer, repoPointer)
		endIfError(err)
	}
	if controllersPointer == nil {
		controllersPointer = controllers.New(servicesPointer)
		endIfError(err)
	}

	ginEngine := router.NewEngine(lambdaContext, controllersPointer)
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		panic("SERVER_PORT not set")
	}

	ginEngine.Run(port)
}

func NewLambdaServer(
	initialized bool,
	ginLambda *ginadapter.GinLambda,
	controllersPointer *controllers.Controller,
	servicesPointer *services.Services,
	validatorPointer *validator.Validate,
	configPointer *config.Config,
	repoPointer *repository.RepositoryManager,
) func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(lambdaContext context.Context, request events.APIGatewayProxyRequest) (_ events.APIGatewayProxyResponse, err error) {
		if initialized {
			return ginLambda.Proxy(request)
		}
		// not initialized
		if configPointer == nil {
			configPointer = config.New()
			endIfError(err)
		}
		if validatorPointer == nil {
			validatorPointer = validator.New()
			endIfError(err)
		}
		if repoPointer == nil {
			repoPointer = repository.New(lambdaContext, configPointer, validatorPointer)
			endIfError(err)
		}
		if servicesPointer == nil {
			servicesPointer, err = services.New(validatorPointer, configPointer, repoPointer)
			endIfError(err)
		}
		if controllersPointer == nil {
			controllersPointer = controllers.New(servicesPointer)
			endIfError(err)
		}

		if !initialized {
			ginEngine := router.NewEngine(lambdaContext, controllersPointer)
			ginLambda = ginadapter.New(ginEngine)
			initialized = true
		}

		return ginLambda.Proxy(request)
	}
}
