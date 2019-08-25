package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Stub dummy route handler func
func Stub(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}
