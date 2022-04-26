package k8s

import (
	"dongzhai/apis/k8s"

	"github.com/gin-gonic/gin"
)

func SecretRouter(g *gin.RouterGroup) {
	secret := g.Group("/secret")

	secret.GET("/", k8s.GetSecrets)
	secret.POST("/", k8s.CreateSecret)
	secret.PATCH("/", k8s.UpdateSecret)
	secret.DELETE("/:id", k8s.DeleteSecretById)
}
