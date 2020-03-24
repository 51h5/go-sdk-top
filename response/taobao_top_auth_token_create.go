package response

type TaobaoTopAuthTokenCreateResponse struct {
    topResponse
    Data taobaoTopAuthTokenCreateResponse `json:"top_auth_token_create_response"`
}

// @see https://open.taobao.com/api.htm?docId=25388&docType=2
type taobaoTopAuthTokenCreateResponse struct {
    TokenResult string `json:"token_result"`
    RequestId   string `json:"request_id"`
}
