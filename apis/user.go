package apis

import (
	"dongzhai/models"
	"dongzhai/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.CreateUser(user); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "add user success", nil, c)
}

func GetUsers(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	users, total, err := service.GetUsers(query)
	if err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	data := gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  users,
	}
	Response(http.StatusOK, "get user success", data, c)
}

func UpdateUser(c *gin.Context) {}

func DeleteUserById(c *gin.Context) {}
