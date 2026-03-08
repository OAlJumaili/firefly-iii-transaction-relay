package web

import (
	"bytes"
	"encoding/json"
	"firefly-iii-transaction-relay/core"
	"firefly-iii-transaction-relay/parser"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{}

func PostTransaction(transaction parser.Transaction, c *gin.Context) {
	payload := map[string]interface{}{
		"transactions": []parser.Transaction{transaction},
	}

	// FOR TESTING PURPOSES ONLY
	// c.JSON(http.StatusOK, gin.H{"message": "Message Parsed successfully", "payload": payload})

	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/transactions", core.FireflyAddress), bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to Compose HTTP Request: %v", err)})
		return
	}

	req.Header.Set("Authorization", "Bearer "+core.FireflyPAT)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)

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

	body, _ := io.ReadAll(resp.Body)
	c.JSON(http.StatusOK, gin.H{
		"message":               "Message Parsed and posted successfully",
		"composed_request":      payload,
		"firefly_response_code": resp.StatusCode,
		"firefly_response":      string(body)})
	return
}
