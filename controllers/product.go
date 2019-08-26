package controllers

// The controller defined below is the product controller,
//
//  Some methods are stubbed out for you to implement them from scratch
// while others may contain one or two bugs
//
// NB: Check the BACKEND CHALLENGE TEMPLATE DOCUMENTATION linked in the readme of this repository
// to see our recommended endpoints, request body/param, and response object for each of these method

import (
	"net/http"

	"github.com/TuringCom/golang_challenge_template/models"
	"github.com/gin-gonic/gin"
)

// GetAllProducts gets all products
func GetAllProducts(c *gin.Context) {
	products := models.GetProducts()
	c.JSON(http.StatusOK, products)
}

// GetProduct gets single product details
func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// SearchProduct search all products
func SearchProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetProductsByCategory get all products in a category
func GetProductsByCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetProductsByDepartment get all products in a department
func GetProductsByDepartment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetAllDepartments get all departments
func GetAllDepartments(c *gin.Context) {
	departments := models.GetDepartments()
	c.JSON(http.StatusOK, departments)
}

// GetDepartment get single department details
func GetDepartment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetAllCategories get all categories
func GetAllCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetCategory get single category details
func GetCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}

// GetDepartmentCategories get all categories in a department
func GetDepartmentCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "NOT IMPLEMENTED"})
}
