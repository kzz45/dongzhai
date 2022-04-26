package k8s

import (
	"dongzhai/apis"
	"dongzhai/models"
	k8s_service "dongzhai/service/k8s"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetClusterNodes(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	cluster_id, ok := c.GetQuery("cluster_id")
	if ok {
		id, _ := strconv.Atoi(cluster_id)
		nodes, total, err := k8s_service.GetClusterNodes(query, id)
		if err != nil {
			apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"page":  query.Page,
			"size":  query.Size,
			"total": total,
			"data":  nodes,
		})
		return
	}
	apis.Response(http.StatusOK, "need cluster_id", nil, c)
}
