package apis

import (
	"dongzhai/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	addr := fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port)
	route := gin.New()

	apis := route.Group("/api")
	clusterRouter(apis)

	route.Run(addr)
}
