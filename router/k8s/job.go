package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func JobsRouter(g *gin.RouterGroup) {
	job := g.Group("/job")

	job.GET("/", k8s.GetJobs)
	job.POST("/", k8s.CreateJob)
}
