package monitor

import (
	"dongzhai/apis/monitor"

	"github.com/gin-gonic/gin"
)

func ServerRouter(g *gin.RouterGroup) {
	server := g.Group("/server")

	server.GET("/", monitor.GetServers)
	server.POST("/", monitor.CreateServer)
	server.PATCH("/", monitor.UpdateServer)
	server.DELETE("/:id", monitor.DeleteServerById)
}
