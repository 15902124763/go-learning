package packet

import (
	"github.com/labstack/echo"
	"go-learning/router/response"
	"net/http"
	"go-learning/service"
	"strconv"
)

// 新增商品
func Add(c echo.Context) (err error) {
	product := new(service.ProductBO)
	if  err = c.Bind(product); err != nil{
		return c.JSONBlob(http.StatusBadRequest, nil)
	}
	// 保存
	service.Save(product)
	return c.JSON(http.StatusOK, response.OK())
}

// 删除商品
func Del(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, response.ParamError)
	}

	i, e := strconv.Atoi(id)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(nil))
	}

	service.Del(i)
	return c.JSON(http.StatusOK, response.OK())
}

// 更新商品
func Update(c echo.Context) (err error) {
	product := new(service.ProductBO)
	if  err = c.Bind(product); err != nil{
		return c.JSONBlob(http.StatusBadRequest, nil)
	}
	// 保存
	service.Update(product)
	return c.JSON(http.StatusOK, response.OK())
}

// 获取商品列表
func GetList(c echo.Context) error {
	param := c.QueryParam("id")

	if param == "" {
		all := service.SelectAll()
		return c.JSON(http.StatusOK, response.Of(all))
	}

	i, e := strconv.Atoi(param)

	if e != nil  {
		return c.JSON(http.StatusInternalServerError, response.Error(nil))
	}
	product := service.Select(i)
	return c.JSON(http.StatusOK, product)
}

