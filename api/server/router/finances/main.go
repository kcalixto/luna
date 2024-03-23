package finances

import (
	"github.com/gin-gonic/gin"
	"github.com/kcalixto/mojo-jojo/api/controllers"
	"github.com/kcalixto/mojo-jojo/api/controllers/viewmodels"
	"github.com/kcalixto/mojo-jojo/api/server/router/routerutils/response"
)

func RegisterFinancesRoutes(router *gin.RouterGroup, crtl *controllers.Controller) {
	r := newFinancesRouter(crtl)

	router.POST("/income/:action", r.HandleIncome)
	router.POST("/expense/:action", r.HandleExpense)
}

type financesRouter struct {
	ctrl *controllers.Controller
}

func newFinancesRouter(crtl *controllers.Controller) *financesRouter {
	return &financesRouter{
		ctrl: crtl,
	}
}

func (r *financesRouter) HandleIncome(ctx *gin.Context) {
	var request viewmodels.IncomePayload
	err := ctx.Bind(&request)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	res, err := r.ctrl.Finances.HandleIncome(request)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Success(ctx, res)
}

func (r *financesRouter) HandleExpense(ctx *gin.Context) {
	var request viewmodels.ExpensePayload
	err := ctx.Bind(&request)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	res, err := r.ctrl.Finances.HandleExpense(request)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Success(ctx, res)
}
