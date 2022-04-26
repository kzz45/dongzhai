package router

import (
	"dongzhai/apis"

	"github.com/gin-gonic/gin"
)

func userGroupRouters(g *gin.RouterGroup) {
	user_group := g.Group("/user_group")

	user_group.GET("/", apis.GetUserGroup)
	user_group.POST("/", apis.CreateUserGroup)
	user_group.PATCH("/", apis.UpdateUserGroup)
	user_group.DELETE("/:id", apis.DeleteUserGroup)
}
