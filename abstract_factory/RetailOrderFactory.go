package abstract_factory

import "design/abstract_factory/impl"

type RetailOrderFactory struct {
}

func (*RetailOrderFactory) CreateOrderFactory() ICreateOrder {
	return &impl.RetailCreateOrder{}
}

func (*RetailOrderFactory) SupplementOrderFactory() ISupplementOrder {
	return &impl.RetailSupplementOrder{}
}