package service

import (
	"gospacex-pengyilong/bff/basic/config"
	"gospacex-pengyilong/bff/handler/request"
	__ "gospacex-pengyilong/srv/proto/goods"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoodsAdd(c *gin.Context) {
	var form request.GoodsAdd
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误",
			"code": 400,
		})
		return
	}
	r, err := config.GoodsClient.GoodsAdd(c, &__.GoodsAddReq{
		Name:  form.Name,
		Price: float32(form.Price),
		Num:   int64(form.Num),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  r.Msg,
		"code": r.Code,
	})
	return
}
