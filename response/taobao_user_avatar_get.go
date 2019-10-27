package response

type TaobaoUserAvatarGetResponse struct {
    topResponse
    Data  taobaoUserAvatarGetResponse `json:"user_avatar_get_response"`
}

// @see https://open.taobao.com/api.htm?source=search&docId=26303&docType=2
type taobaoUserAvatarGetResponse struct {
    Avatar string `json:"avatar"`
}
