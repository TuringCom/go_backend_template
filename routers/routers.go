package routers

import (
	"github.com/TuringCom/golang_challenge_template/controllers"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

// router.GET("/v1/images/:path1", GetHandler)           //      /v1/images/detail
// router.GET("/v1/images/:path1/:path2", GetHandler)    //      /v1/images/<id>/history

// func GetHandler(c *gin.Context) {
// path1 := c.Param("path1")
// path2 := c.Param("path2")

// if path1 == "detail" && path2 == "" {
// 		Detail(c)
// } else if path1 != "" && path2 == "history" {
// 		imageId := path1
// 		History(c, imageId)
// } else {
// 		HandleHttpError(c, NewHttpError(404, "Page not found"))
// }
// }

func getHandlerForRoute(c *gin.Context) {

}

// SetupRouter initialize routes and handlers
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/attributes/", controllers.Stub)
	// r.GET("/attributes/:attribute_id", controllers.Stub)
	r.GET("/attributes/values/:attribute_id", controllers.Stub)
	r.GET("/attributes/inProduct/:product_id", controllers.Stub)

	r.POST("/customers", controllers.Stub)
	r.POST("/customers/login", controllers.Stub)

	r.GET("/customer", controllers.Stub)
	r.PUT("/customer", controllers.Stub)
	r.PUT("/customer/address", controllers.Stub)
	r.PUT("/customer/creditCard", controllers.Stub)

	r.GET("/products", controllers.Stub)
	// r.GET("/products/:product_id", controllers.Stub)
	r.GET("/products/search", controllers.Stub)
	r.GET("/products/inCategory/:category_id", controllers.Stub)
	r.GET("/products/inDepartment/:department_id", controllers.Stub)

	r.GET("/departments", controllers.Stub)
	r.GET("/departments/:department_id", controllers.Stub)

	r.GET("/categories", controllers.Stub)
	// r.GET("/categories/:category_id", controllers.Stub)
	r.GET("/categories/inDepartment/:department_id", controllers.Stub)

	r.GET("/shipping/regions", controllers.Stub)
	r.GET("/shipping/regions/:shipping_region_id", controllers.Stub)

	r.GET("/shoppingcart/generateUniqueId", controllers.Stub)
	r.POST("/shoppingcart/add", controllers.Stub)
	// r.GET("/shoppingcart/:cart_id", controllers.Stub)
	r.PUT("/shoppingcart/update/:item_id", controllers.Stub)
	r.DELETE("/shoppingcart/empty/:cart_id", controllers.Stub)
	r.DELETE("/shoppingcart/removeProduct/:item_id", controllers.Stub)

	r.POST("/orders", controllers.Stub)
	r.GET("/orders/inCustomer", controllers.Stub)
	// r.GET("/orders/:order_id", controllers.Stub)

	r.POST("/stripe/charge", controllers.Stub)

	r.GET("/tax", controllers.Stub)
	r.GET("/tax/:tax_id", controllers.Stub)

	return r
}
