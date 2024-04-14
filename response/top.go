package response

// TOP协议返回 - 错误结果
type topErrorResponse struct {
	Code      int    `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	SubCode   string `json:"sub_code,omitempty"`
	SubMsg    string `json:"sub_msg,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

// TOP协议返回 - 共享基类
type topResponse struct {
	Error topErrorResponse `json:"error_response"`
}

func (r *topResponse) Success() bool {
	return r.Error.Code == 0 && r.Error.SubCode == ""
}

func (r *topResponse) Fix() {}
