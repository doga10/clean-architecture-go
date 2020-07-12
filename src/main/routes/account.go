package routes

import (
	"github.com/doga10/clean-architecture-go/src/main/adapters"
	"github.com/doga10/clean-architecture-go/src/main/factories/controllers/account/add-account"
	load_account_by_id "github.com/doga10/clean-architecture-go/src/main/factories/controllers/account/load-account-by-id"
	"github.com/gin-gonic/gin"
)

func AccountRouter(router gin.IRouter) {
	router.POST("/v1/accounts", adapters.AdaptRouter(add_account.MakeAddAccountController()))
	router.GET("/v1/accounts/:id", adapters.AdaptRouter(load_account_by_id.MakeLoadAccountByIdController()))
}