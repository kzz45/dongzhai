package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func IngressRouters(g *gin.RouterGroup) {
	ingress := g.Group("/ingress")

	ingress.GET("/", k8s.GetIngresses)
	ingress.POST("/", k8s.CreateIngress)
	ingress.PATCH("/", k8s.UpdateIngress)
	ingress.DELETE("/:id", k8s.DeleteIngressById)
}
