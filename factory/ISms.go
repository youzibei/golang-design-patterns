package factory

// 定义SMS接口
type ISms interface {
	// SMS发送方法
	Send(params SendVO) error
}
