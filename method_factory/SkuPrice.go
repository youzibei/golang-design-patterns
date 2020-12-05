package method_factory

import "fmt"

type SkuPrice struct {
	*PriceBase
}

// 通过SKU，获取SKU价格
func (p *SkuPrice) GetPrice() float64 {
	sku := p.sku
	fmt.Printf("通过sku：%s 通获取价格", sku)

	return 1.00
}