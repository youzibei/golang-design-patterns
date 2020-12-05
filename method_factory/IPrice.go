package method_factory

// 价格接口
type IPrice interface {
	// SKU
	SetSku(string)
	// 客户
	SetCompany(string)
	// 商店
	SetStore(string)
	// 单价
	GetPrice() float64
}