package router

import (
	"dongzhai/apis"

	"github.com/gin-gonic/gin"
)

func roleRouter(g *gin.RouterGroup) {
	role := g.Group("/role")

	role.GET("/", apis.GetRoles)
	role.POST("/", apis.CreateRole)
	role.PATCH("/", apis.UpdateRole)
	role.DELETE("/:id", apis.DeleteRoleById)
}
