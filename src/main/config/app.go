package config

import (
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()

	Routes(router)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
