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

func CreateConfigMap(c *gin.Context) {
	var configmap k8s_model.ConfigMap
	if err := c.Bind(&configmap); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.CreateConfigMap(configmap); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "create configmap success", nil, c)
}

func GetConfigMaps(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	configmaps, total, err := k8s_service.GetConfigMaps(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  configmaps,
	})
}

func UpdateConfigMap(c *gin.Context) {
	var configmap k8s_model.ConfigMap
	if err := c.Bind(&configmap); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.UpdateConfigMap(configmap); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "update configmap success", configmap, c)
}

func DeleteConfigMapById(c *gin.Context) {
	id := c.Param("id")
	configmap_id, err := strconv.Atoi(id)
	if err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = k8s_service.DeleteConfigMapById(configmap_id); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "delete configmap success", nil, c)
}
