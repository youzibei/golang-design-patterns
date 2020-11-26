package factory

import "testing"

func TestSend(t *testing.T) {
	fa := SmsSendFactory{}
	sendVO := SendVO{}
	sendVO.mobile = "13166299999"
	sendVO.msg = "你的短信验证是：1311"
	fa.Create("10002").Send(sendVO)
}
