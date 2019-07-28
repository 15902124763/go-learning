package router

import (
	"github.com/labstack/echo"
	"go-learning/router/api"
)

func RouterInit(e *echo.Echo) {
	api.ApiRouters(e)
}
