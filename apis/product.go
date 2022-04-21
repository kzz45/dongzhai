package apis

import (
	"dongzhai/models"
	"dongzhai/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.CreateProduct(product); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "add product success", nil, c)
}

func GetProducts(c *gin.Context) {
	query := &models.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	products, total, err := service.GetProducts(query)
	if err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	data := gin.H{
		"page":  query.Page,
		"size":  query.Size,
		"total": total,
		"data":  products,
	}
	Response(http.StatusOK, "get products success", data, c)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err := service.UpdateProduct(product); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "update product success", product, c)
}

func DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		Response(http.StatusBadRequest, fmt.Sprintf("%s", err), nil, c)
		return
	}
	if err = service.DeleteProductById(product_id); err != nil {
		Response(http.StatusInternalServerError, fmt.Sprintf("%s", err), nil, c)
		return
	}
	Response(http.StatusOK, "delete product success", nil, c)
}
