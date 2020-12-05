package method_factory

// 价格公用类
type PriceBase struct {
	sku string
	company string
	store string
}

func (p *PriceBase) SetSku(sku string) {
	p.sku = sku
}

func (p *PriceBase) SetCompany(company string) {
	p.company = company
}

func (p *PriceBase) SetStore(store string) {
	p.store = store
}