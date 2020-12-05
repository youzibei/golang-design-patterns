package abstract_factory

// 补单接口
type ISupplementOrder interface {
	SupplementOrder(params map[string]interface{}) map[string]interface{}
}