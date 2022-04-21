package apis

import (
	"dongzhai/models"
	"dongzhai/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCloud(c *gin.Context) {
	var cloud models.Cloud
	if err := c.Bind(&cloud); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.CreateCloud(cloud); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "add cloud success", nil, c)
}

func GetClouds(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	clouds, total, err := service.GetClouds(query)
	if err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	data := gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  clouds,
	}
	Response(http.StatusOK, "get clouds success", data, c)
}

func UpdateCloud(c *gin.Context) {
	var cloud models.Cloud
	if err := c.Bind(&cloud); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.UpdateCloud(cloud); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "update cloud success", cloud, c)
}

func DeleteCloudById(c *gin.Context) {
	id := c.Param("id")
	cloud_id, err := strconv.Atoi(id)
	if err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = service.DeleteCloudById(cloud_id); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "delete cloud success", nil, c)
}
