package apis

import (
	"dongzhai/models"
	"dongzhai/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUserGroup(c *gin.Context) {
	var group models.UserGroup
	if err := c.Bind(&group); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	// group.Users
	if err := service.CreateUserGroup(group); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "add user group success", nil, c)
}

func GetUserGroup(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	ugs, total, err := service.GetUserGroup(query)
	if err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  ugs,
	})
}

func UpdateUserGroup(c *gin.Context) {
	var group models.UserGroup
	if err := c.Bind(&group); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.UpdateUserGroup(group); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "update user_group success", group, c)
}

func DeleteUserGroup(c *gin.Context) {
	id := c.Param("id")
	group_id, err := strconv.Atoi(id)
	if err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = service.DeleteUserGroup(group_id); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "delete user_group success", nil, c)
}
