package request

import (
    "51h5.com/sdk/top/internal/utils"
    "net/url"
)

type TaobaoMixnickGetRequest struct {
    topRequest
    Nick string `json:"nick"`
}

func (r *TaobaoMixnickGetRequest) Method() string {
    return "taobao.mixnick.get"
}

func (r *TaobaoMixnickGetRequest) Values() url.Values {
    return url.Values{
        "nick": {r.Nick},
    }
}

func (r *TaobaoMixnickGetRequest) Check() (code uint, err error) {
    code, err = utils.CheckNotEmpty("nick", r.Nick)
    if err != nil {
        return
    }

    return
}