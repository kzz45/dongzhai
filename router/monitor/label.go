package monitor

import (
	"dongzhai/apis/monitor"

	"github.com/gin-gonic/gin"
)

func LabelRouter(g *gin.RouterGroup) {
	label := g.Group("/label")

	label.GET("/name", monitor.GetLabelName)
	label.POST("/name", monitor.CreateLabelName)
	label.GET("/value", monitor.GetLabelValue)
	label.POST("/value", monitor.CreateLabelValue)
}
