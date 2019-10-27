package response

type TaobaoMixnickGetResponse struct {
    topResponse
    Data taobaoMixnickGetResponse `json:"mixnick_get_response"`
}

// @see https://open.taobao.com/api.htm?source=search&docId=26303&docType=2
type taobaoMixnickGetResponse struct {
    Nick string `json:"nick"`
}
