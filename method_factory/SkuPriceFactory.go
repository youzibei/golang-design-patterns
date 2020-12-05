package method_factory

type SkuPriceFactory struct {
}

// Sku价格工厂
func (p *SkuPriceFactory) Create() IPrice {
	return &SkuPrice{
		PriceBase: &PriceBase{},
	}
}