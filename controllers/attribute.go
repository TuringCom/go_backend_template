package controllers

// The controller defined below is the attribute controller,
//
//  Some methods are stubbed out for you to implement them from scratch
// while others may contain one or two bugs
//
// NB: Check the BACKEND CHALLENGE TEMPLATE DOCUMENTATION linked in the readme of this repository
// to see our recommended endpoints, request body/param, and response object for each of these method

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllAttributes get all attributes
func GetAllAttributes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetSingleAttribute get single attribute by id
func GetSingleAttribute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetAttributeValues get attribute values
func GetAttributeValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetProductAttributes get attributes for product
func GetProductAttributes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}
