package sms

// SendRequest 发送短信请求参数
type SendRequest struct {
	CommonRequest
	PhoneNumbers    string `json:"PhoneNumbers"`
	SignName        string `json:"SignName"`
	TemplateCode    string `json:"TemplateCode"`
	OutID           string `json:"OutId"`           // 外部流水扩展字段
	SmsUpExtendCode string `json:"SmsUpExtendCode"` // 上行扩展码
	TemplateParam   string `json:"TemplateParam"`   // 短信模板变量对应的实际值，JSON格式
}

// SendBatchRequest 批量发送短信请求参数
type SendBatchRequest struct {
	CommonRequest
	PhoneNumberJSON     string `json:"PhoneNumberJson"`
	SignNameJSON        string `json:"SignNameJson"`
	TemplateCode        string `json:"TemplateCode"`
	SmsUpExtendCodeJSON string `json:"SmsUpExtendCodeJson"`
	TemplateParamJSON   string `json:"TemplateParamJson"`
}

// SendResponse 发送短信返回值
type SendResponse struct {
	Response
	BizID string `json:"BizId"`
}

// Send 发送短信
func (s *Sms) Send(req *SendRequest) (*SendResponse, error) {
	// req.Sign()
	req.Action = "SendSms"
	resp := &SendResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// SendBatch 批量发送短信
func (s *Sms) SendBatch(req *SendBatchRequest) (*SendResponse, error) {
	req.Action = "SendBatchSms"
	resp := &SendResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
