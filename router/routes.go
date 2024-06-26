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
		v1.POST("/product/:id", controllers.CreateProduct)
		v1.DELETE("/product", controllers.DeleteProduct)
		v1.PUT("/product", controllers.UpdateProduct)
		v1.GET("/products", controllers.GetAllProducts)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
