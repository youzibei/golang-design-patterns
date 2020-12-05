package method_factory

type CompanyPriceFactory struct {
}

// 客户定制价工厂
func (p *CompanyPriceFactory) Create() IPrice {
	return &CompanyPrice{
		PriceBase: &PriceBase{},
	}
}