package main

import (
	"firefly-iii-transaction-relay/core"
	"firefly-iii-transaction-relay/parser"
	"firefly-iii-transaction-relay/web"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Request struct {
	AuthKey string `json:"auth_key"`
	Message string `json:"message"`
}

func postTransaction(c *gin.Context) {
	var req Request

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.AuthKey != core.AuthKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth key"})
		return
	}

	msg := req.Message
	msg = strings.ReplaceAll(msg, "\n", " ")
	transaction := parser.ParseMessage(msg)
	web.PostTransaction(transaction, c)
}

func main() {
	core.InitEnv()
	core.InitRegex()
	router := gin.Default()
	router.POST("/transaction", postTransaction)
	router.Run(core.ListenAddress)
}
