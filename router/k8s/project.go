package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func ProjectRouter(g *gin.RouterGroup) {
	project := g.Group("/project")

	project.GET("/", k8s.GetProjects)
	project.POST("/", k8s.CreateProject)
	project.PATCH("/", k8s.UpdateProject)
	project.DELETE("/:id", k8s.DeleteProjectById)
}
