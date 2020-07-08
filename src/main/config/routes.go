package config

import (
	"github.com/doga10/clean-architecture-go/src/main/routes"
	"github.com/gin-gonic/gin"
)

func Routes(router gin.IRouter) {
	routes.AccountRouter(router)
}
