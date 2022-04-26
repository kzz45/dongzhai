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

func GetDeployments(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	project_id, ok := c.GetQuery("project_id")
	id, _ := strconv.Atoi(project_id)
	if ok {
		deploys, total, err := k8s_service.GetDeploymentWithAppId(query, id)
		if err != nil {
			apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"page":  query.Page,
			"size":  query.Size,
			"total": total,
			"data":  deploys,
		})
		return
	}
	deploys, total, err := k8s_service.GetDeployment(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  deploys,
	})
}

func CreateDeployment(c *gin.Context) {
	var deploy k8s_model.Deployment
	if err := c.Bind(&deploy); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.CreateDeployment(deploy); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "create deployment success", nil, c)
}

func UpdateDeployment(c *gin.Context) {
	var deploy k8s_model.Deployment
	if err := c.Bind(&deploy); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.UpdateDeployment(deploy); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "update deploy success", deploy, c)
}
