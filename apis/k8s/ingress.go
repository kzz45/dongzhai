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

func CreateIngress(c *gin.Context) {
	var ingress k8s_model.Ingress
	if err := c.Bind(&ingress); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.CreateIngress(ingress); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "create ingress success", nil, c)
}

func GetIngresses(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	ingresses, total, err := k8s_service.GetIngresses(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  ingresses,
	})
}

func UpdateIngress(c *gin.Context) {
	var ingress k8s_model.Ingress
	if err := c.Bind(&ingress); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.UpdateIngress(ingress); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "update ingress success", ingress, c)
}

func DeleteIngressById(c *gin.Context) {
	id := c.Param("id")
	ingress_id, err := strconv.Atoi(id)
	if err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = k8s_service.DeleteIngressById(ingress_id); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "delete ingress success", nil, c)
}
