package main

import (
	"bytes"
	"encoding/json"
	"firefly-iii-transaction-relay/initialize"
	"firefly-iii-transaction-relay/parser"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{}

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
	msg = strings.ReplaceAll(msg, "\n", " ")
	transaction := parser.ParseMessage(msg)

	payload := map[string]interface{}{
		"transactions": []parser.Transaction{transaction},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "Message Parsed successfully", "transaction": transaction})

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/transactions", initialize.FireflyAddress), bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to Compose HTTP Request: %v", err)})
		return
	}
	request.Header.Set("Authorization", "Bearer "+initialize.FireflyPAT)
	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to post transaction: %v", err)})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to post transaction: Firefly API returned %v:%v", resp.StatusCode, string(body))})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message Parsed and posted successfully"})
	return

}

func main() {
	initialize.InitEnv()
	initialize.InitRegex()
	router := gin.Default()
	router.POST("/transaction", postTransaction)
	router.Run(initialize.ListenAddress)
}
