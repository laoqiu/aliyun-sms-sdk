package sms

import (
	"reflect"
	"strings"
)

// DefaultCommonRequest 固定公共参数
var DefaultCommonRequest = map[string]string{
	"SignatureMethod":  "HMAC-SHA1",
	"SignatureVersion": "1.0",
	"Version":          "2017-05-25",
}

// Response 返回参数
type Response struct {
	RequestID string `json:"RequestId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
}

// Request 请求参数接口
type Request struct {
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	SignatureNonce   string `json:"SignatureNonce"`
	Version          string `json:"Version"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	Action           string `json:"Action"` // 每个接口对应的Action
}

// SpecialURLEncode 返回特殊URL编码处理后的值
func (req *Request) SpecialURLEncode(value string) string {
	return ""
}

// Sign 返回签名值
func (req *Request) Sign(accessSecret string, encodeStr string) string {
	return ""
}

// ToString 返回拼接字符串
func (req *Request) ToString() string {
	return ""
}

// StructToMap struct转map
func StructToMap(i interface{}) map[string]interface{} {
	values := make(map[string]interface{})
	if i != nil {
		iVal := reflect.ValueOf(i).Elem()
		tp := iVal.Type()
		for i := 0; i < iVal.NumField(); i++ {
			tag := tp.Field(i).Tag.Get("json")
			if len(tag) > 0 {
				name := strings.Split(tag, ",")[0]
				if name != "-" {
					values[name] = iVal.Field(i).Interface()
				}
			}
		}
	}
	return values
}
