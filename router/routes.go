package router

import (
	"github.com/WillianIsami/go_api/controllers"
	docs "github.com/WillianIsami/go_api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRouter(router *gin.Engine) {
	controllers.InitializeControllers()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/product", controllers.GetProduct)
		v1.POST("/product", controllers.CreateProduct)
		v1.DELETE("/product", controllers.DeleteProduct)
		v1.PUT("/product", controllers.UpdateProduct)
		v1.GET("/products", controllers.GetAllProducts)

		v1.GET("/category", controllers.GetCategory)
		v1.POST("/category", controllers.CreateCategory)
		v1.DELETE("/category", controllers.DeleteCategory)
		v1.PUT("/category", controllers.UpdateCategory)
		v1.GET("/categories", controllers.GetAllCategory)

		v1.GET("/order", controllers.GetOrder)
		v1.POST("/order", controllers.CreateOrder)
		v1.GET("/orders", controllers.GetAllOrders)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
