package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func ServiceRouter(g *gin.RouterGroup) {
	service := g.Group("/service")

	service.GET("/", k8s.GetServices)
	service.POST("/", k8s.CreateService)
	service.GET("/ports/", k8s.GetServicePorts)
}
