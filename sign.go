package sms

type AddSmsSignRequest struct {
	CommonRequest
	SignName string
	// 签名来源。其中：
	// 0：企事业单位的全称或简称。
	// 1：工信部备案网站的全称或简称。
	// 2：APP应用的全称或简称。
	// 3：公众号或小程序的全称或简称。
	// 4：电商平台店铺名的全称或简称。
	// 5：商标名的全称或简称
	SignSource int
	Remark     string
}

type SmsSignResponse struct {
	SignName  string `json:"SignName"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	Code      string `json:"Code"`
}

type ModifySmsSignRequest struct {
	CommonRequest
	SignName   string
	SignSource int
	Remark     string
}

type DeleteSmsSignRequest struct {
	CommonRequest
	SignName string
}

type QuerySmsSignRequest struct {
	CommonRequest
	SignName string
}

type QuerySmsSignResponse struct {
	SignName   string `json:"SignName"`
	Message    string `json:"Message"`
	RequestID  string `json:"RequestId"`
	Code       string `json:"Code"`
	CreateDate string `json:"CreateDate"`
	Reason     string `json:"Reason"`
	SignStatus int    `json:"SignStatus"`
}

// AddSmsSign 申请短信签名
func (s *Sms) AddSmsSign(req *AddSmsSignRequest) (*SmsSignResponse, error) {
	// req.Sign()
	req.Action = "AddSmsSign"
	resp := &SmsSignResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ModifySmsSign 修改未审核通过的短信签名，并重新提交审核
func (s *Sms) ModifySmsSign(req *DeleteSmsSignRequest) (*SmsSignResponse, error) {
	// req.Sign()
	req.Action = "ModifySmsSign"
	resp := &SmsSignResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteSmsSign 删除短信签名
func (s *Sms) DeleteSmsSign(req *DeleteSmsSignRequest) (*SmsSignResponse, error) {
	// req.Sign()
	req.Action = "DeleteSmsSign"
	resp := &SmsSignResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QuerySmsSign 查询短信签名申请状态
func (s *Sms) QuerySmsSign(req *QuerySmsSignRequest) (*QuerySmsSignResponse, error) {
	// req.Sign()
	req.Action = "QuerySmsSign"
	resp := &QuerySmsSignResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
