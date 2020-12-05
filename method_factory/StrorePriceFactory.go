package method_factory

type StorePriceFactory struct {
}

// 商店工厂
func (p *StorePriceFactory) Create() IPrice {
	return &StorePrice{
		PriceBase: &PriceBase{},
	}
}