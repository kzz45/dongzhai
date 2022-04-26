package k8s

import (
	"dongzhai/apis"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	k8s_service "dongzhai/service/k8s"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context) {
	var svc k8s_model.Service
	if err := c.Bind(&svc); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.CreateService(svc); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "create service success", nil, c)
}

func GetServices(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	services, total, err := k8s_service.GetServices(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  services,
	})
}

func GetServicePorts(c *gin.Context) {
	service_id, ok := c.GetQuery("service_id")
	if ok {
		id, _ := strconv.Atoi(service_id)
		svc_ports, err := k8s_service.GetServicePorts(id)
		if err != nil {
			apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
			return
		}
		apis.Response(http.StatusOK, "get service port success", svc_ports, c)
		return
	}
	apis.Response(http.StatusBadRequest, "service_id is null", nil, c)
}
