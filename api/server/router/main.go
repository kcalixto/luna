package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kcalixto/mojo-jojo/api/controllers"
	"github.com/kcalixto/mojo-jojo/api/server/router/finances"
)

func NewEngine(
	ctrl *controllers.Controller,
) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	versionGroup := engine.Group("/v1")

	// Register routes
	financesGroup := versionGroup.Group("/finances")
	finances.RegisterFinancesRoutes(financesGroup, ctrl)

	return engine
}
