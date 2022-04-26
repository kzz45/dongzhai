package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func DeploymentRouters(g *gin.RouterGroup) {
	deploy := g.Group("/deployment")

	deploy.GET("/", k8s.GetDeployments)     // 获取deploy列表
	deploy.POST("/", k8s.CreateDeployment)  // 新增一个deploy
	deploy.PATCH("/", k8s.UpdateDeployment) //
}
