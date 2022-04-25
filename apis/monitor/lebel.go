package monitor

import (
	"dongzhai/apis"
	"dongzhai/models"
	"dongzhai/models/monitor"
	service "dongzhai/service/monitor"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLabelName(c *gin.Context) {
	var label monitor.Label
	if err := c.Bind(&label); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.CreateLabelName(label); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "add label success", nil, c)
}

func GetLabelName(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	label_names, total, err := service.GetLabelName(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  label_names,
	})
}

func CreateLabelValue(c *gin.Context) {
	var label_value monitor.LabelValue
	if err := c.Bind(&label_value); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.CreateLabelValue(label_value); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "add label_value success", nil, c)
}

func GetLabelValue(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	label_values, total, err := service.GetLabelValue(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  label_values,
	})
}
