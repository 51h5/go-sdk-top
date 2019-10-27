package top

import "time"

const (
    GATEWAY_URL         = "http://gw.api.taobao.com/router/rest"
    FORMAT_TYPE_JSON    = "json"
    API_VERSION         = "2.0"
    PARTNER_ID          = "top-sdk-go-20191025"
    TIME_LAYOUT         = "2006-01-02 15:04:05"
    KEY_SIGN            = "sign"
    HTTP_CLIENT_TIMEOUT = 30 * time.Second
    // KEY_RESPONSE_SUFFIX = "_response"
    // KEY_ERROR           = "error_response"
)
