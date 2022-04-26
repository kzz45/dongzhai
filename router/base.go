package router

import (
	"dongzhai/apis"
	"dongzhai/config"
	"dongzhai/router/k8s"
	"dongzhai/router/monitor"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRouter() {
	addr := fmt.Sprintf("%s:%d", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port)
	if config.GlobalConfig.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.Default()

	v1_base := route.Group("/api/v1")
	v1_base.POST("/login", apis.UserLogin)
	v1_base.POST("/logout", apis.UserLogout)

	{
		userRouter(v1_base)
		roleRouter(v1_base)
		cloudRouter(v1_base)
		productRouter(v1_base)

		k8s.JobsRouter(v1_base)
		k8s.SecretRouter(v1_base)
		k8s.ProjectRouter(v1_base)
		k8s.ClusterRouter(v1_base)
		k8s.ServiceRouter(v1_base)
		k8s.ConfigMapRouter(v1_base)
	}

	v1_monitor := route.Group("/api/v1/monitor")
	{
		monitor.LabelRouter(v1_monitor)
		monitor.ServerRouter(v1_monitor)
	}

	// v1_k8s := route.Group("/api/v1/k8s/")
	// {
	// 	k8s.ClusterRouter(v1_k8s)
	// }
	server_name := config.GlobalConfig.Server.Name
	logrus.Infof("%s running at: %s", server_name, addr)
	route.Run(addr)
}
