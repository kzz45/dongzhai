package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func PodRouters(g *gin.RouterGroup) {
	pod := g.Group("/pod")

	pod.GET("/", k8s.GetPods)
}
