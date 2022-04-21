package router

import (
	"dongzhai/apis"

	"github.com/gin-gonic/gin"
)

func userRouter(g *gin.RouterGroup) {
	user := g.Group("/user")

	user.GET("/", apis.GetUsers)
	user.POST("/", apis.CreateUser)
	user.PATCH("/", apis.UpdateUser)
	user.DELETE("/:id", apis.DeleteUserById)
}
