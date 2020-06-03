package sms

// QuerySendDetailsRequest 查看短信发送记录和发送状态请求参数
type QuerySendDetailsRequest struct {
	CommonRequest
	CurrentPage int    `json:"CurrentPage"`
	PageSize    int    `json:"PageSize"`
	PhoneNumber string `json:"PhoneNumber"`
	SendDate    string `json:"SendDate"`
	BizID       string `json:"BizId"`
}

// QuerySendDetailsResponse 短信发送记录返回值
type QuerySendDetailsResponse struct {
	Response
	SmsSendDetailDTOs []SmsSendDetailDTOs `json:"SmsSendDetailDTOs"`
	TotalCount        int                 `json:"TotalCount"`
}

// SmsSendDetailDTOs 短信发送记录
type SmsSendDetailDTOs struct {
	SmsSendDetailDTO struct {
		SendDate     string `json:"SendDate"`
		OutID        int    `json:"OutId"`
		SendStatus   int    `json:"SendStatus"`
		ReceiveDate  string `json:"ReceiveDate"`
		ErrCode      string `json:"ErrCode"`
		TemplateCode string `json:"TemplateCode"`
		Content      string `json:"Content"`
		PhoneNum     int64  `json:"PhoneNum"`
	} `json:"SmsSendDetailDTO"`
}

// QuerySendDetails 查看短信发送记录和发送状态
func (s *Sms) QuerySendDetails(req *QuerySendDetailsRequest) (*QuerySendDetailsResponse, error) {
	req.Action = "QuerySendDetails"
	resp := &QuerySendDetailsResponse{}
	if err := s.Fetch(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
