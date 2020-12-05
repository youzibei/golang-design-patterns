package method_factory

import "fmt"

type StorePrice struct {
	*PriceBase
}

// 通过SKU与商店，获取商店定制价格
func (p *StorePrice) GetPrice() float64 {
	sku := p.sku
	store := p.store
	fmt.Printf("通过sku与商店：%s，%s，获取价格", sku, store)
	
	return 1.00
}