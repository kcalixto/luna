package finances

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kcalixto/mojo-jojo/api/controllers"
	"github.com/kcalixto/mojo-jojo/api/controllers/viewmodels"
	responseUtils "github.com/kcalixto/mojo-jojo/api/server/router/routerutils/response"
)

func RegisterFinancesRoutes(lambdaContext context.Context, router *gin.RouterGroup, ctrl *controllers.Controller) {
	r := newFinancesRouter(lambdaContext, ctrl)

	router.POST("/income/:action", r.HandleIncome)
	router.POST("/expense/:action", r.HandleExpense)
}

type financesRouter struct {
	lambdaContext context.Context
	ctrl          *controllers.Controller
}

func newFinancesRouter(lambdaContext context.Context, ctrl *controllers.Controller) *financesRouter {
	return &financesRouter{lambdaContext, ctrl}
}

func (r *financesRouter) HandleIncome(ctx *gin.Context) {
	var request viewmodels.IncomeRequestPayload
	err := ctx.Bind(&request)
	if err != nil {
		responseUtils.BadRequest(ctx, err.Error())
		return
	}
	var requestUri viewmodels.IncomeRequestUri
	err = ctx.ShouldBindUri(&requestUri)
	if err != nil {
		responseUtils.BadRequest(ctx, err.Error())
		return
	}

	response, err := r.ctrl.Finances.HandleIncome(r.lambdaContext, requestUri, request)
	if err != nil {
		responseUtils.InternalServerError(ctx, err.Error())
		return
	}

	responseUtils.Success(ctx, response)
}

func (r *financesRouter) HandleExpense(ctx *gin.Context) {
	var request viewmodels.ExpenseRequestPayload
	err := ctx.Bind(&request)
	if err != nil {
		responseUtils.BadRequest(ctx, err.Error())
		return
	}
	var requestUri viewmodels.ExpenseRequestUri
	err = ctx.ShouldBindUri(&requestUri)
	if err != nil {
		responseUtils.BadRequest(ctx, err.Error())
		return
	}

	response, err := r.ctrl.Finances.HandleExpense(r.lambdaContext, requestUri, request)
	if err != nil {
		responseUtils.InternalServerError(ctx, err.Error())
		return
	}

	responseUtils.Success(ctx, response)
}
