package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome route handler func
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Turing E-commerce shop api, your goal is to implement the missing code or fix the bugs inside this project"})
}
