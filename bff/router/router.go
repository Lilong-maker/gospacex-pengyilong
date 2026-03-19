package router

import (
	"gospacex-pengyilong/bff/api"
	"gospacex-pengyilong/bff/handler/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	r.POST("GoodsAdd", service.GoodsAdd)
	r.POST("/notify/pay", api.NotifyPay)
	return r
}
