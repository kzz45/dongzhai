package router

import (
	"dongzhai/config"
	"dongzhai/router/monitor"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	addr := fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port)
	if config.GlobalConfig.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.Default()

	v1_monitor := route.Group("/api/v1/monitor")
	{
		monitor.ServerRouter(v1_monitor)
	}

	route.Run(addr)
}
