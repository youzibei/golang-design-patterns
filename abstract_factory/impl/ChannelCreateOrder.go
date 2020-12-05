package impl

import "fmt"

// 渠道创建订单
type ChannelCreateOrder struct {
}

// 实现订单
func (r *ChannelCreateOrder) CreateOrder(params map[string]interface{}) map[string]interface{} {
	fmt.Println("创建渠道订单")
	response := make(map[string]interface{})
	return response
}