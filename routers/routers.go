package routers

import (
	"net/http"

	"github.com/TuringCom/golang_challenge_template/controllers"
	"github.com/gin-gonic/gin"
)

// returns the appropriate handler func based on the matched route wildcard
// simple workaround for issues like https://github.com/gin-gonic/gin/issues/205
// since we are using httprouter from Gin
func getHandlerFromWildCard(m map[string]gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		wildcard1 := c.Param("wildcard1")

		if handler := m[wildcard1]; handler != nil {
			handler(c)
		} else if wildcard1 != "" {
			wildcard1 = "id_param"
			handler := m[wildcard1]
			handler(c)
		} else {
			// not found
			c.String(http.StatusNotFound, "404 page not found")
		}
	}
}

var attributesRouteHandlerFuncMap = map[string]gin.HandlerFunc{
	"values":    controllers.GetAttributeValues,   // "/attributes/values/:attribute_id"
	"inProduct": controllers.GetProductAttributes, // "/attributes/inProduct/:product_id"
}

var productsRouteHandlerFuncMap = map[string]gin.HandlerFunc{
	"id_param":     controllers.GetProduct,              // "/products/:product_id"
	"search":       controllers.SearchProduct,           // "/products/search"
	"inCategory":   controllers.GetProductsByCategory,   // "/products/inCategory/:category_id"
	"inDepartment": controllers.GetProductsByDepartment, // "/products/inDepartment/:department_id"
}

var shoppingCartRouteHandlerFuncMap = map[string]gin.HandlerFunc{
	"generateUniqueId": controllers.GenerateUniqueCart, // "/shoppingcart/generateUniqueId"
	"id_param":         controllers.GetCart,            // "/shoppingcart/:cart_id"
}

var ordersRouteHandlerFuncMap = map[string]gin.HandlerFunc{
	"inCustomer": controllers.GetCustomerOrders, // "/orders/inCustomer"
	"id_param":   controllers.GetOrderSummary,   // "/orders/:order_id"
}

// SetupRouter initialize routes and handlers
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.Welcome)

	r.GET("/attributes/", controllers.GetAllAttributes)
	// This matches "/attributes/:attribute_id"
	r.GET("/attributes/:wildcard1", controllers.GetSingleAttribute)
	// This matches both "/attributes/values/:attribute_id" and "/attributes/inProduct/:product_id"
	r.GET("/attributes/:wildcard1/:wildcard2", getHandlerFromWildCard(attributesRouteHandlerFuncMap))

	r.POST("/customers", controllers.Create)
	r.POST("/customers/login", controllers.Login)

	r.GET("/customer", controllers.GetCustomerProfile)
	r.PUT("/customer", controllers.UpdateCustomerProfile)
	r.PUT("/customer/address", controllers.UpdateCustomerAddress)
	r.PUT("/customer/creditCard", controllers.UpdateCreditCard)

	r.GET("/products", controllers.GetAllProducts)
	// This matches "/products/:product_id" and "/products/search"
	r.GET("/products/:wildcard1", getHandlerFromWildCard(productsRouteHandlerFuncMap))
	// This matches "/products/inCategory/:category_id" and "/products/inDepartment/:department_id"
	r.GET("/products/:wildcard1/:wildcard2", getHandlerFromWildCard(productsRouteHandlerFuncMap))

	r.GET("/departments", controllers.GetAllDepartments)
	r.GET("/departments/:department_id", controllers.GetDepartment)

	r.GET("/categories", controllers.GetAllCategories)
	// This matches // "/categories/:category_id"
	r.GET("/categories/:wildcard1", controllers.GetCategory)
	// This matches "/categories/inDepartment/:department_id"
	r.GET("/categories/:wildcard1/:department_id", controllers.GetDepartmentCategories)

	r.GET("/shipping/regions", controllers.GetShippingRegions)
	r.GET("/shipping/regions/:shipping_region_id", controllers.GetShippingType)

	r.POST("/shoppingcart/add", controllers.AddItemToCart)
	// This matches "/shoppingcart/generateUniqueId" and "/shoppingcart/:cart_id"
	r.GET("/shoppingcart/:wildcard1", getHandlerFromWildCard(shoppingCartRouteHandlerFuncMap))
	r.PUT("/shoppingcart/update/:item_id", controllers.UpdateCartItem)
	r.DELETE("/shoppingcart/empty/:cart_id", controllers.EmptyCart)
	r.DELETE("/shoppingcart/removeProduct/:item_id", controllers.RemoveItemFromCart)

	r.POST("/orders", controllers.CreateOrder)
	// This matches // "/orders/inCustomer" and "/orders/:order_id"
	r.GET("/orders/:wildcard1", getHandlerFromWildCard(ordersRouteHandlerFuncMap))

	r.POST("/stripe/charge", controllers.ProcessStripePayment)

	r.GET("/tax", controllers.GetAllTax)
	r.GET("/tax/:tax_id", controllers.GetSingleTax)

	return r
}
