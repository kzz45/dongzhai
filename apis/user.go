package apis

import (
	"dongzhai/models"
	"dongzhai/service"
	"fmt"
	"net/http"
	"strconv"

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
	id, ok := c.GetQuery("id")
	if ok {
		user_id, err := strconv.Atoi(id)
		if err != nil {
			Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
			return
		}
		user, err := service.GetUserById(user_id)
		if err != nil {
			Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
			return
		}
		Response(http.StatusOK, "get user success", user, c)
		return
	}
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

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.UpdateUser(user); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "update user success", user, c)
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = service.DeleteUserById(user_id); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "delete user success", nil, c)
}

func UserLogin(c *gin.Context) {
	var user_login models.UserLogin
	if err := c.Bind(&user_login); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	user_resp, err := service.UserLogin(user_login)
	if err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "login success", user_resp, c)
}

func UserLogout(c *gin.Context) {
	Response(http.StatusOK, "logout success", nil, c)
}
