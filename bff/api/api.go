package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gospacex-pengyilong/srv/basic/config"
	_ "gospacex-pengyilong/srv/basic/inits"
	"gospacex-pengyilong/srv/handler/model"
)

func NotifyPay(c *gin.Context) {
	c.Request.ParseForm()
	fmt.Println("1111", c.Request.PostForm)
	TradeStatus := c.PostForm("trade_status")
	if TradeStatus != "TRADE_SUCCESS" {
		return
	}
	outTradeNo := c.PostForm("out_trade_no")
	if outTradeNo == "" {
		return
	}

	tx := config.DB.Begin()

	var order model.Order
	err := tx.Where("order_no = ?", outTradeNo).First(&order).Error
	if err != nil {
		tx.Rollback()
		return
	}
	order.PayStatus = 2
	err = tx.Save(&order).Error
	if err != nil {
		tx.Rollback()
		return
	}

	var orderItems []model.OrderItem
	err = tx.Where("id=?", order.ID).Find(&orderItems).Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, items := range orderItems {
		var goods model.Goods
		err = tx.Where("id=?", items.GoodsID).First(&goods).Error
		if err != nil {
			tx.Rollback()
			return
		}
		goods.Num -= items.Num
		err = tx.Save(&goods).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
}
