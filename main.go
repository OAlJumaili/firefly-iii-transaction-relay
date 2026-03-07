package main

import (
	"bytes"
	"encoding/json"
	"firefly-iii-transaction-relay/initialize"
	"firefly-iii-transaction-relay/parser"
	"fmt"
	"net/http"

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

	if req.AuthKey != initialize.AuthKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth key"})
		return
	}

	msg := req.Message
	transaction := parser.ParseMessage(msg)

	jsonData, err := json.Marshal(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(initialize.FireflyAddress, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to post transaction: %v", resp.StatusCode)})
		return
	}
}

func main() {
	initialize.InitEnv()
	initialize.InitRegex()
	router := gin.Default()
	router.POST("/transaction", postTransaction)
	router.Run(initialize.ListenAddress)
}
