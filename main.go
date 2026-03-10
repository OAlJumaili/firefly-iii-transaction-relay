package main

import (
	"firefly-iii-transaction-relay/core"
	"firefly-iii-transaction-relay/parser"
	"firefly-iii-transaction-relay/web"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Request struct {
	AuthKey         string `json:"auth_key"`
	Message         string `json:"message"`
	VerificationKey string `json:"verification_key"`
}

func postTransaction(c *gin.Context) {
	var req Request

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.AuthKey != core.AuthKey || req.VerificationKey != core.VerificationKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "One or more authentication keys are invalid."})
		return
	}

	msg := req.Message
	msg = strings.ReplaceAll(msg, "\n", " ")
	transaction, blacklisted := parser.ParseMessage(msg)
	if blacklisted {
		c.JSON(http.StatusOK, gin.H{"message": "Blacklisted Keyword Detected, Skipping."})
		return
	}
	web.PostTransaction(transaction, c)
}

func main() {
	core.InitEnv()
	core.InitRegex()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/transaction", postTransaction)
	fmt.Println(`HTTP Listening and serving traffic on`, core.ListenAddress)
	router.Run(core.ListenAddress)
}
