package routes

import (
	"github.com/doga10/clean-architecture-go/src/main/adapters"
	"github.com/doga10/clean-architecture-go/src/main/factories/controllers/account"
	"github.com/gin-gonic/gin"
)

func AccountRouter(router gin.IRouter) {
	router.POST("/v1/accounts", adapters.AdaptRouter(account.MakeAddAccountController()))
}