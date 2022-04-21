package router

import (
	"dongzhai/apis"

	"github.com/gin-gonic/gin"
)

func productRouter(g *gin.RouterGroup) {
	product := g.Group("/product")

	product.GET("/", apis.GetProducts)
	product.POST("/", apis.CreateProduct)
	product.PATCH("/", apis.UpdateProduct)
	product.DELETE("/:id", apis.DeleteProductById)
}
