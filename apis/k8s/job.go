package k8s

import (
	"dongzhai/apis"
	"dongzhai/models"
	k8s_model "dongzhai/models/k8s"
	k8s_service "dongzhai/service/k8s"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	var job k8s_model.Job
	if err := c.Bind(&job); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.CreateJob(job); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "create job success", nil, c)
}

func GetJobs(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	jobs, total, err := k8s_service.GetJobs(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  jobs,
	})
}
