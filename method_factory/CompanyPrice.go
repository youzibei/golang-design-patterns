package method_factory

import "fmt"

type CompanyPrice struct {
	*PriceBase
}

// 通过sku与客户，获取客户定制价
func (p *CompanyPrice) GetPrice() float64 {
	sku := p.sku
	company := p.company
	fmt.Printf("通过sku与客户：%s，%s，获取价格", sku, company)

	return 1.00
}