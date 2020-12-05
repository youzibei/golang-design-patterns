package abstract_factory

import "fmt"

type OrderService struct {
}

func (*OrderService) createOrder(params map[string]interface{}) interface{} {
	channelId := fmt.Sprintf("%v", params["channId"])
	if channelId == "1" {
		fmt.Printf("创建零售订单")
	} else if channelId == "2" {
		fmt.Printf("创建渠道订单")
	}

	return "订单创建成功"
}
