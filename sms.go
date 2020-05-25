package sms

import (
	"io/ioutil"
	"net/http"
)

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

// Fetch 发送请求
func (s *Sms) Fetch(req Request) ([]byte, error) {
	q := req.ToString(s.accessSecret)
	// 发送http请求
	res, err := http.Get("http://dysmsapi.aliyuncs.com/?" + q)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
