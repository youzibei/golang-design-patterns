package method_factory

// 价格工厂
type PriceFactory interface {
	Create() IPrice
}