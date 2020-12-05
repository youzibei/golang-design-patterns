package abstract_factory

type IOrderFactory interface {
	CreateOrderFactory() ICreateOrder
	SupplementOrderFactory() ISupplementOrder
}