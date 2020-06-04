package sms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
)

// DefaultCommonRequest 固定公共参数
var DefaultCommonRequest = map[string]string{
	"SignatureMethod":  "HMAC-SHA1",
	"SignatureVersion": "1.0",
	"Version":          "2017-05-25",
}

// Request 请求类接口
type Request interface {
	ActionName() string
}

// Response 返回参数
type Response struct {
	RequestID string `json:"RequestId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
}

// CommonRequest 请求参数接口
type CommonRequest struct {
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	SignatureNonce   string `json:"SignatureNonce"`
	Version          string `json:"Version"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	Action           string `json:"Action"` // 每个接口对应的Action
}

func (req *CommonRequest) ActionName() string {
	return req.Action
}

func (req *CommonRequest) SetDefaultParam() error {
	b, _ := json.Marshal(DefaultCommonRequest)
	return json.Unmarshal(b, req)
}

// ToString 返回拼接字符串
func ToString(req Request, secret string) string {
	v := structToValues(req)
	sortedQueryString := specialURLEncode(v.Encode())
	stringToSign := "GET&%2F&" + url.QueryEscape(sortedQueryString) // 注意再次urlencode
	sign := sign(secret+"&", stringToSign)
	v.Set("Signature", sign)
	q := specialURLEncode(v.Encode())
	return q
}

// structToValues struct转url.Values
func structToValues(i interface{}) url.Values {
	// values := make(map[string]interface{})
	v := url.Values{}
	if i != nil {
		iVal := reflect.ValueOf(i).Elem()
		tp := iVal.Type()
		for i := 0; i < iVal.NumField(); i++ {
			tag := tp.Field(i).Tag.Get("json")
			if len(tag) > 0 {
				name := strings.Split(tag, ",")[0]
				if name != "-" {
					v.Set(name, iVal.Field(i).String())
				}
			}
		}
	}
	return v
}

// specialURLEncode 返回特殊URL编码处理后的值
func specialURLEncode(value string) string {
	value = strings.Replace(value, "+", "%20", -1)
	value = strings.Replace(value, "*", "%2A", -1)
	value = strings.Replace(value, "%7E", "~", -1)
	return value
}

// sign 返回签名值
func sign(accessSecret string, encodeStr string) string {
	hash := hmac.New(sha1.New, []byte(accessSecret))
	hash.Write([]byte(encodeStr))
	hstring := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return hstring
}
