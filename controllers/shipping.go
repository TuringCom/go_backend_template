package controllers

// The controller defined below is the shipping controller,
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

// GetShippingRegions get all shipping regions
func GetShippingRegions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetShippingType get all shipping type for a shipping region
func GetShippingType(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}
