package monitor

import (
	"dongzhai/apis"
	"dongzhai/models"
	"dongzhai/models/monitor"
	service "dongzhai/service/monitor"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateServer(c *gin.Context) {
	var server monitor.Server
	if err := c.Bind(&server); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.CreateServer(server); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "add server success", nil, c)
}

func GetServers(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	servers, total, err := service.GetServers(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	data := gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  servers,
	}
	apis.Response(http.StatusOK, "get server success", data, c)
}

func UpdateServer(c *gin.Context) {
	var server monitor.Server
	if err := c.Bind(&server); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.UpdateServer(server); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "update server success", server, c)
}

func DeleteServerById(c *gin.Context) {
	id := c.Param("id")
	cloud_id, err := strconv.Atoi(id)
	if err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = service.DeleteServerById(cloud_id); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "delete cloud success", nil, c)
}
