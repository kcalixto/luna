package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kcalixto/luna/api/controllers"
	"github.com/kcalixto/luna/api/server/router/account"
	"github.com/kcalixto/luna/api/server/router/finances"
)

func NewEngine(
	lambdaContext context.Context,
	ctrl *controllers.Controller,
) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	versionGroup := engine.Group("/v1")

	// Register routes
	financesGroup := versionGroup.Group("/finances")
	finances.RegisterFinancesRoutes(lambdaContext, financesGroup, ctrl)

	// Register routes
	accountGroup := versionGroup.Group("/account")
	account.RegisterAccountRoutes(lambdaContext, accountGroup, ctrl)

	return engine
}
