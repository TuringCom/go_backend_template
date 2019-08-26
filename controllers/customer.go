package controllers

// The controller defined below is the customer controller,
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

// Create adds new customer
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// Login log in customer
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetCustomerProfile gets the logged in customer's profile
func GetCustomerProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// UpdateCustomerProfile updates the logged in customer's profile such as
// name, email, password, day_phone, eve_phone and mob_phone
func UpdateCustomerProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// UpdateCustomerAddress updates the logged in customer's address such as
// address_1, address_2, city, region, postal_code, country and shipping_region_id
func UpdateCustomerAddress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// UpdateCreditCard updates the logged in customer's credit card
func UpdateCreditCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}
