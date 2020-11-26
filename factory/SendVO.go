package factory

// SMS发送实体
type SendVO struct {
	// 手机号
	mobile string
	// 发送的消息
	msg string
}
