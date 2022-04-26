package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func RegistryRouters(g *gin.RouterGroup) {
	registry := g.Group("/registry")

	registry.GET("/", k8s.GetRegistries)            // 获取仓库列表
	registry.POST("/", k8s.CreateRegistry)          // 新增一个仓库
	registry.PATCH("/", k8s.UpdateRegistry)         // 更新某个仓库
	registry.DELETE("/:id", k8s.DeleteRegistryById) // 删除某个仓库
}
