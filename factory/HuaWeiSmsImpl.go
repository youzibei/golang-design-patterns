package factory

import "fmt"

type HuaWeiSmsImpl struct {
}

// 实现ISms接口Send方法， 调用华为云发送
func (s *HuaWeiSmsImpl) Send(param SendVO) error {
	fmt.Println("调用华为云SMS接口，发送短信", param.mobile, param.msg)
	return nil
}
