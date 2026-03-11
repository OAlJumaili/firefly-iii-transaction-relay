package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionRequest struct {
	VerificationKey string `json:"verification_key"`
	Message         string `json:"message"`
}
type AuthRequest struct {
	AuthKey string `json:"auth_key"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Relay is up and running!"})
}
