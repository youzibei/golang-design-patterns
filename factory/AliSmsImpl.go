package factory

import "fmt"

type AliSmsImpl struct {
}

// 实现ISms接口Send方法，调用阿里云发送
func (s *AliSmsImpl) Send(param SendVO) error {
	fmt.Println("调用阿里云SMS接口，发送短信", param.mobile, param.msg)
	return nil
}
