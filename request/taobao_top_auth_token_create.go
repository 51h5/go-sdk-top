package request

import (
	"net/url"

	"github.com/51h5/go-sdk-top/internal/utils"
)

type TaobaoTopAuthTokenCreateRequest struct {
	topRequest
	Code string `json:"code"`
	Uuid string `json:"uuid"`
}

func (r *TaobaoTopAuthTokenCreateRequest) Method() string {
	return "taobao.top.auth.token.create"
}

func (r *TaobaoTopAuthTokenCreateRequest) Values() url.Values {
	return url.Values{
		"code": {r.Code},
		"uuid": {r.Uuid},
	}
}

func (r *TaobaoTopAuthTokenCreateRequest) Check() (code uint, err error) {
	code, err = utils.CheckNotEmpty("code", r.Code)
	if err != nil {
		return
	}

	return
}
