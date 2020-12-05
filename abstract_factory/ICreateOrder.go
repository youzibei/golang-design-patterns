package abstract_factory

// 创建订单接口
type ICreateOrder interface {
	CreateOrder(params map[string]interface{}) map[string]interface{}
}