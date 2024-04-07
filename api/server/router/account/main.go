package account

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kcalixto/mojo-jojo/api/controllers"
	"github.com/kcalixto/mojo-jojo/api/server/router/routerutils/response"
)

func RegisterAccountRoutes(lambdaContext context.Context, router *gin.RouterGroup, ctrl *controllers.Controller) {
	r := newAccountRouter(lambdaContext, ctrl)

	router.GET("/transaction/list", r.HandleTransactionList)
}

type accountRouter struct {
	lambdaContext context.Context
	ctrl          *controllers.Controller
}

func newAccountRouter(lambdaContext context.Context, ctrl *controllers.Controller) *accountRouter {
	return &accountRouter{lambdaContext, ctrl}
}

func (r *accountRouter) HandleTransactionList(ctx *gin.Context) {
	response, err := r.ctrl.Account.HandleListTransactions(r.lambdaContext)
	if err != nil {
		responseUtils.InternalServerError(ctx, err.Error())
		return
	}

	responseUtils.Success(ctx, response)
}
