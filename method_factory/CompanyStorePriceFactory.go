package method_factory

type CompanyStorePriceFactory struct {
}

// 客户商店定制价格工厂
func (p *CompanyStorePriceFactory) Create() IPrice {
	return &CompanyStorePrice{
		PriceBase: &PriceBase{},
	}
}