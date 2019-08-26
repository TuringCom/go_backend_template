package controllers

// The controller defined below is the shopping cart controller,
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

// GenerateUniqueCart generate random unique id for cart identifier
func GenerateUniqueCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// AddItemToCart add item to existing cart with cart id
func AddItemToCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetCart get all items in a shopping cart using cart id
func GetCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// UpdateCartItem update quantity of an item in a shopping cart
func UpdateCartItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// EmptyCart remove all items from a shopping cart
func EmptyCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// RemoveItemFromCart remove a specific item from a shopping cart
func RemoveItemFromCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// CreateOrder create order for all items in a shopping cart
func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetCustomerOrders get all orders placed by a customer
func GetCustomerOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetOrderSummary get details of items in an order
func GetOrderSummary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// ProcessStripePayment checkout order and process stripe payment
func ProcessStripePayment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}
