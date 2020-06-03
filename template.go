package sms

type AddSmsTemplateRequest struct {
	CommonRequest
	// 短信类型。其中：
	// 0：验证码。
	// 1：短信通知。
	// 2：推广短信。
	// 3：国际/港澳台消息。
	TemplateType int
	// 模板名称，长度为1~30个字符
	TemplateName string
	// 模板内容，长度为1~500个字符
	TemplateContent string
	// 短信模板申请说明
	Remark string
}

type SmsTemplateResponse struct {
	Message      string `json:"Message"`
	RequestID    string `json:"RequestId"`
	TemplateCode string `json:"TemplateCode"`
	Code         string `json:"Code"`
}

type DeleteSmsTemplateRequest struct {
	CommonRequest
	// 短信模板CODE
	TemplateCode string
}

type ModifySmsTemplateRequest struct {
	CommonRequest
	TemplateType int
	// 模板名称，长度为1~30个字符
	TemplateName string
	// 模板内容，长度为1~500个字符
	TemplateContent string
	// 短信模板申请说明
	Remark string
}

type QuerySmsTemplateRequest struct {
	CommonRequest
	// 短信模板CODE
	TemplateCode string
}

type QuerySmsTemplateResponse struct {
	TemplateType    int    `json:"TemplateType"`
	TemplateContent string `json:"TemplateContent"`
	TemplateName    string `json:"TemplateName"`
	Message         string `json:"Message"`
	RequestID       string `json:"RequestId"`
	TemplateCode    string `json:"TemplateCode"`
	CreateDate      string `json:"CreateDate"`
	Code            string `json:"Code"`
	TemplateStatus  int    `json:"TemplateStatus"`
	Reason          string `json:"Reason"`
}

// AddSmsTemplate 申请短信模板
func (s *Sms) AddSmsTemplate(req *AddSmsTemplateRequest) (*SmsTemplateResponse, error) {
	// req.Sign()
	req.Action = "AddSmsTemplate"
	resp := &SmsTemplateResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteSmsTemplate 删除短信模板
func (s *Sms) DeleteSmsTemplate(req *DeleteSmsTemplateRequest) (*SmsTemplateResponse, error) {
	// req.Sign()
	req.Action = "DeleteSmsTemplate"
	resp := &SmsTemplateResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ModifySmsTemplate 修改未通过审核的短信模板
func (s *Sms) ModifySmsTemplate(req *ModifySmsTemplateRequest) (*SmsTemplateResponse, error) {
	// req.Sign()
	req.Action = "ModifySmsTemplate"
	resp := &SmsTemplateResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QuerySmsTemplate 修改未通过审核的短信模板
func (s *Sms) QuerySmsTemplate(req *QuerySmsTemplateRequest) (*QuerySmsTemplateResponse, error) {
	// req.Sign()
	req.Action = "QuerySmsTemplate"
	resp := &QuerySmsTemplateResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
