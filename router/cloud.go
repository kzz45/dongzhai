package router

import (
	"dongzhai/apis"

	"github.com/gin-gonic/gin"
)

func cloudRouter(g *gin.RouterGroup) {
	cloud := g.Group("/cloud")

	cloud.GET("/", apis.GetClouds)
	cloud.POST("/", apis.CreateCloud)
	cloud.PATCH("/", apis.UpdateCloud)
	cloud.DELETE("/:id", apis.DeleteCloudById)
}
