/*
Package api api路由器
*/
package api

import (
	"github.com/labstack/echo"
	"go-learning/router/api/mudles/packet"
)

/**
路由
*/
func ApiRouters(e *echo.Echo) {
	packet.ProductRouter(e)
}
