package apis

import (
	"dongzhai/controllers"

	"github.com/gin-gonic/gin"
)

func clusterRouter(g *gin.RouterGroup) {
	cluster := g.Group("/cluster")

	cluster.GET("/", controllers.GetClusters)
	cluster.POST("/", controllers.CreateCluster)
}
