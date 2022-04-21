package apis

import (
	"dongzhai/models"

	"github.com/gin-gonic/gin"
)

func Response(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(code, models.Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
