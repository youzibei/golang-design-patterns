package abstract_factory

import (
	"testing"
)

func TestOrder(t *testing.T)  {
	// 创建渠道订单
	var factory IOrderFactory
	factory = &ChannelOrderFactory{}
	m := make(map[string]interface{})
	m["price"] = "2.33"
	m["sku"] = "S343"
	factory.CreateOrderFactory().CreateOrder(m)
}