package method_factory

import "fmt"

type CompanyStorePrice struct {
	*PriceBase
}

// 通过sku与客户+商店，获取客户在不同商店的定制价
func (p *CompanyStorePrice) GetPrice() float64 {
	sku := p.sku
	company := p.company
	store := p.store
	fmt.Printf("通过sku与客户+商店：%s,%s,%s, 获取价格", sku, company, store)

	return 1.00
}