package main

import (
	"fmt"
	"github.com/labstack/echo"
	"go-learning/conf"
	"go-learning/log"
	"go-learning/router"
	"go-learning/router/response"
	"net/http"
)

func main() {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	router.RouterInit(e)
	e.Logger.Info(e.Start(conf.Conf.Server.Port))
}

//路由错误页面配置
func customHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Internal)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = fmt.Sprintf("%s : %v", http.StatusText(code), err)
	}
	if _, ok := msg.(string); ok {
		msg = c.JSON(http.StatusInternalServerError, response.Error(&response.GlobalErrorCode{ErrorCode: code, Error: fmt.Sprint(msg)}))
	}
	// Send response
	if !c.Response().Committed {
		// Issue #608
		if c.Request().Method == "HEAD" {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			log.Error(err)
		}
	}
}
