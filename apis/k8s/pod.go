package k8s

import (
	"dongzhai/apis"
	"dongzhai/models"
	k8s_service "dongzhai/service/k8s"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPods(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	pods, total, err := k8s_service.GetPods(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  pods,
	})
}
