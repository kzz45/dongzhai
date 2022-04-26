package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func ConfigMapRouter(g *gin.RouterGroup) {
	configmap := g.Group("/configmap")

	configmap.GET("/", k8s.GetConfigMaps)
	configmap.POST("/", k8s.CreateConfigMap)
	configmap.PATCH("/", k8s.UpdateConfigMap)
	configmap.DELETE("/:id", k8s.DeleteConfigMapById)
}
