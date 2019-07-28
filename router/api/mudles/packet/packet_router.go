package packet

import (
	"github.com/labstack/echo"
)

// 路由器
func ProductRouter(e *echo.Echo) {
	// 增
	e.POST("/v1/go/learn/add", Add)
	// 删
	e.POST("/v1/go/learn/del", Del)
	// 改
	e.POST("/v1/go/learn/update", Update)
	// 查
	e.GET("/v1/go/learn/get/list", GetList)
}
