package controllers

import (
	"dongzhai/client"
	"dongzhai/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCluster(c *gin.Context) {
	var cluster models.Cluster
	if err := c.Bind(&cluster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  fmt.Sprintf("%s", err),
		})
		return
	}
	if err := client.DBClient.Create(&cluster).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "create cluster success",
	})
}

func GetClusters(c *gin.Context) {
	var clusters []models.Cluster
	if err := client.DBClient.Find(&clusters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "get all cluster success",
		"data": clusters,
	})
}
