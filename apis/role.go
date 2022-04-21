package apis

import (
	"dongzhai/models"
	"dongzhai/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.CreateRole(role); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "add role success", nil, c)
}

func GetRoles(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	roles, total, err := service.GetRoles(query)
	if err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	data := gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  roles,
	}
	Response(http.StatusOK, "get roles success", data, c)
}

func UpdateRole(c *gin.Context) {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.UpdateRole(role); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "update role success", role, c)
}

func DeleteRoleById(c *gin.Context) {
	id := c.Param("id")
	role_id, err := strconv.Atoi(id)
	if err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = service.DeleteRoleById(role_id); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "delete role success", nil, c)
}
