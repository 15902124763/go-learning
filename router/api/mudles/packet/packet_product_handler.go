package packet

import (
	"github.com/labstack/echo"
	"go-learning/router/response"
	"net/http"
)

// 新增商品
func Add(c echo.Context) (err error) {
	product := new(Product)
	if  err = c.Bind(product); err != nil{
		return c.JSONBlob(http.StatusBadRequest, nil)
	}
	println(product)
	return c.JSON(http.StatusOK, response.Of(product))
}

// 删除商品
func Del(c echo.Context) error {
	return c.JSON(http.StatusOK, response.OK())
}

// 更新商品
func Update(c echo.Context) error {
	return c.JSON(http.StatusOK, response.OK())
}

// 获取商品列表
func GetList(c echo.Context) error {
	return c.JSON(http.StatusOK, response.OK())
}

type (
	Product struct {
		productName string `json:"productName" form:"productName" query:"productName"`
		productUrl string `json:"productUrl" form:"productUrl" query:"productUrl"`
	}
)
