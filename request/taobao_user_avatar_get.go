package request

import (
    "51h5.com/sdk/top/internal/utils"
    "net/url"
)

type TaobaoUserAvatarGetRequest struct {
    topRequest
    Nick string `json:"nick"`
}

func (r *TaobaoUserAvatarGetRequest) Method() string {
    return "taobao.user.avatar.get"
}

func (r *TaobaoUserAvatarGetRequest) Values() url.Values {
    return url.Values{
        "nick": {r.Nick},
    }
}

func (r *TaobaoUserAvatarGetRequest) Check() (code uint, err error) {
    code, err = utils.CheckNotEmpty("nick", r.Nick)
    if err != nil {
        return
    }

    return
}
