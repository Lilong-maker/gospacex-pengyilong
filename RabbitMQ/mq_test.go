package RabbitMQ

import (
	"testing"
)

func TestSendMsg(t *testing.T) {
	Send("test_queue", "hello mq")
}

func TestSubscribeMsg(t *testing.T) {
	Subscribe("test_queue", func(msg string) {
		t.Log("消费：", msg)
	})
}

func TestStockDeduct(t *testing.T) {
	SendStockDeductMsg("1001", 1)
}
