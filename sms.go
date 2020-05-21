package sms

// DefaultEndpoint 短信服务地址
var DefaultEndpoint = "dysmsapi.aliyuncs.com"

// Sms 短信服务
type Sms struct {
	accessKey    string
	accessSecret string
	endpoint     string
}

// NewSms 返回sms实例
func NewSms(key string) *Sms {
	return &Sms{
		accessKey: key,
		endpoint:  DefaultEndpoint,
	}
}
