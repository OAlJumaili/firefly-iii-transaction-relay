package main

import (
	"firefly-iii-transaction-relay/core"
	"firefly-iii-transaction-relay/web"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	core.InitEnv()
	core.InitRegex()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/health", web.Health)
	router.POST("/auth", web.CreateSession)
	router.POST("/transaction", web.AuthSession, web.ProcessTransaction)

	fmt.Println(`HTTP Listening and serving traffic on`, core.ListenAddress)
	router.Run(core.ListenAddress)
}
