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

func CreateRegistry(c *gin.Context) {
	var registry k8s_model.Registry
	if err := c.Bind(&registry); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.CreateRegistry(registry); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "create registry success", nil, c)
}

func GetRegistries(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	registries, total, err := k8s_service.GetRegistries(query)
	if err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  registries,
	})
}

func UpdateRegistry(c *gin.Context) {
	var registry k8s_model.Registry
	if err := c.Bind(&registry); err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := k8s_service.UpdateRegistry(registry); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "update registry success", registry, c)
}

func DeleteRegistryById(c *gin.Context) {
	id := c.Param("id")
	registry_id, err := strconv.Atoi(id)
	if err != nil {
		apis.Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = k8s_service.DeleteRegistryById(registry_id); err != nil {
		apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	apis.Response(http.StatusOK, "delete registry success", nil, c)
}

func ImageSearch(c *gin.Context) {
	word, _ := c.GetQuery("word")
	id, ok := c.GetQuery("registry_id")
	if ok {
		registry_id, _ := strconv.Atoi(id)
		registry, err := k8s_service.GetRegistryById(registry_id)
		if err != nil {
			apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
			return
		}
		image_list := k8s_service.ImageSearch(registry, word)
		apis.Response(http.StatusOK, "get image from harbor success", image_list, c)
		return
	}
	image_list := k8s_service.SearchDockerHub(word)
	apis.Response(http.StatusOK, "get image from dockerhub success", image_list, c)
}

func GetImageTags(c *gin.Context) {
	image, _ := c.GetQuery("image")
	id, ok := c.GetQuery("registry_id")
	if ok {
		registry_id, _ := strconv.Atoi(id)
		registry, err := k8s_service.GetRegistryById(registry_id)
		if err != nil {
			apis.Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
			return
		}
		tag_list := k8s_service.GetImageTags(registry, image)
		apis.Response(http.StatusOK, "get image tags from harbor success", tag_list, c)
		return
	}
	tag_list := k8s_service.GetTagInDockerHub(image)
	apis.Response(http.StatusOK, "get image tags from dockerhub success", tag_list, c)
}
