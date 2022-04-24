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

func CreateCluster(c *gin.Context) {
	var cluster k8s_model.Cluster
	if err := c.Bind(&cluster); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.CreateCluster(cluster); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "add cluster success", nil, c)
}

func GetCluster(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	clusters, total, err := k8s_service.GetClusters(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	data := gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  clusters,
	}
	apis.Response(http.StatusOK, "get clusters success", data, c)
}

func UpdateCluster(c *gin.Context) {
	var cluster k8s_model.Cluster
	if err := c.Bind(&cluster); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.UpdateCluster(cluster); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "update cluster success", cluster, c)
}

func DeleteCluster(c *gin.Context) {
	id := c.Param("id")
	cluster_id, err := strconv.Atoi(id)
	if err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = k8s_service.DeleteClusterById(cluster_id); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "delete cluster success", nil, c)
}
