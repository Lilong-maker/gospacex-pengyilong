package main

import (
	"gospacex-pengyilong/RabbitMQ"
	_ "gospacex-pengyilong/srv/basic/inits"
)

func main() {
	RabbitMQ.SendStockDeductMsg("樱桃", 333)
}
