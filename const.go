package top

import "time"

const (
    kGatewayUrl        = "http://gw.api.taobao.com/router/rest"
    kFormat            = "json"
    kApiVersion        = "2.0"
    kPartnerId         = "top-sdk-go-20191025"
    kTimeFormat        = "2006-01-02 15:04:05"
    kContentType       = "application/x-www-form-urlencoded"
    kKeySign           = "sign"
    kHttpClientTimeout = 30 * time.Second
)
