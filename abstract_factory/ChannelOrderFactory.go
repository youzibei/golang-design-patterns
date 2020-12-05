package abstract_factory

import "design/abstract_factory/impl"

type ChannelOrderFactory struct {
}

func (*ChannelOrderFactory) CreateOrderFactory() ICreateOrder {
	return &impl.ChannelCreateOrder{}
}

func (*ChannelOrderFactory) SupplementOrderFactory() ISupplementOrder {
	return &impl.ChannelSupplementOrder{}
}