package factory

// SMS发送工厂
type SmsSendFactory struct {
}

// 通过工厂方法发送SMS
func (s *SmsSendFactory) Create(smsCode string) ISms {
	if smsCode == "10001" {
		return &HuaWeiSmsImpl{}
	} else if smsCode == "10002" {
		return &AliSmsImpl{}
	} else {
		panic("短信服务商不存在")
	}
}
