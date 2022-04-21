package routes

import (
	"sctrans/controllers"
	"sctrans/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var r *gin.Engine = gin.Default()

func SetupRouter() *gin.Engine {
	r.Use(middlewares.ErrorMiddleware())

	customerController := controllers.CustomerController{}
	r.GET("/customers", customerController.GetCustomers)
	r.POST("/customers", customerController.CreateCustomer)
	r.POST("/customers/:id", customerController.AddAddressCustomer)

	productController := controllers.ProductController{}
	r.GET("/products", productController.GetProducts)
	r.POST("/products", productController.CreateProduct)

	paymentMethodController := controllers.PaymentMethodController{}
	r.GET("/payment-methods", paymentMethodController.GetPaymentMethods)
	r.POST("/payment-methods", paymentMethodController.CreatePaymentMethod)

	orderMethodController := controllers.OrderController{}
	r.GET("/orders", orderMethodController.GetOrders)
	r.POST("/orders", orderMethodController.CreateOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
