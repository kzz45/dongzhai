package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func ClusterRouter(g *gin.RouterGroup) {
	cluster := g.Group("/cluster")

	cluster.GET("/", k8s.GetCluster)
	cluster.POST("/", k8s.CreateCluster)
	cluster.PATCH("/", k8s.UpdateCluster)
	cluster.DELETE("/:id", k8s.DeleteCluster)
}
