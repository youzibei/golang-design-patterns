package method_factory

import (
	"fmt"
	"testing"
)

func TestPrice(t *testing.T)  {
	factory := CompanyStorePriceFactory{}
	op := factory.Create()
	op.SetCompany("公司ID")
	op.SetSku("Sku")
	op.SetStore("商店ID")
	price := op.GetPrice()
	fmt.Println("\n获取到的价格是：", price)
}
