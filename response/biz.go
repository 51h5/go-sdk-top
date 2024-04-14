package response

// 业务返回
type bizResponse struct {
	RequestId    string `json:"request_id,omitempty"`
	TraceId      string `json:"trace_id,omitempty"`
	BizSuccess   bool   `json:"biz_success,omitempty"`
	BizErrorCode string `json:"biz_error_code,omitempty"`
	BizErrorMsg  string `json:"biz_error_msg,omitempty"`
}

func (r *bizResponse) Success() bool {
	return r.BizSuccess && r.BizErrorCode == "" && r.BizErrorCode == ""
}

func (r *bizResponse) Fix() {}
