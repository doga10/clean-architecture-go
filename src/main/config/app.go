package config

import (
	"github.com/gin-gonic/gin"
	"os"
)

func StartApp() {
	router := gin.Default()
	gin.SetMode(os.Getenv("APP_ENV"))
	Routes(router)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
