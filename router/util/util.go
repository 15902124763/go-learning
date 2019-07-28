package util

import (
	"fmt"
	"github.com/labstack/echo"
	"gitlub.handeson.com/HundredClouds/echo-rest/cache"
	"gitlub.handeson.com/HundredClouds/echo-rest/middleware"
	"gitlub.handeson.com/HundredClouds/echo-rest/model"
	"strconv"
	"time"
)

func XquarkContext(c echo.Context) *middleware.XquarkContext {
	return c.(*middleware.XquarkContext)
}

func CurrUser(c echo.Context) *model.User {
	return XquarkContext(c).CurrUser()
}

func CacheStore(c echo.Context) cache.Store {
	return XquarkContext(c).CacheStore()
}

func NowDateFormatString() string {
	year, month, day := time.Now().Date()
	return fmt.Sprint("", year, "-", int(month), "-", day)
}

func NowDateString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func NowDateInt() int64 {
	return time.Now().UnixNano() / 1000000
}

// 秒为单位
func NextDaySecondsZeroBetween() int {
	now := time.Now()
	tm2 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1)
	return int(tm2.Sub(now).Seconds())
}
