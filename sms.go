package sms

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DefaultEndpoint 短信服务地址
var DefaultEndpoint = "https://dysmsapi.aliyuncs.com"

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

// SetEndpoint 设置新的短信服务地址
func (s *Sms) SetEndpoint(endpoint string) {
	s.endpoint = endpoint
}

// Fetch 发送请求
func (s *Sms) Fetch(req Request, resp interface{}) error {
	q := ToString(req, s.accessSecret)
	// client不验证https证书
	c := http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	// 发送请求
	res, err := c.Get(s.endpoint + "/?" + q)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}
	return nil
}
