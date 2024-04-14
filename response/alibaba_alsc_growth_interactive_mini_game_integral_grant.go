package response

type AlibabaAlscGrowthInteractiveMiniGameIntegralGrantResponse struct {
	topResponse
	Data alibabaAlscGrowthInteractiveMiniGameIntegralGrantResponse `json:"alibaba_alsc_growth_interactive_mini_game_integral_grant_response"` // 积分发放返回对象
}

// 积分发放返回对象
// @see https://open.taobao.com/api.htm?docId=63518&docType=2
type alibabaAlscGrowthInteractiveMiniGameIntegralGrantResponse struct {
	bizResponse
	AccountValue   int64 `json:"account_value,omitempty"`    // 帐户值
	RealGrantValue int64 `json:"real_grant_value,omitempty"` // 实际发放值
}

func (r *AlibabaAlscGrowthInteractiveMiniGameIntegralGrantResponse) Success() bool {
	return r.topResponse.Success() && r.Data.bizResponse.Success()
}
