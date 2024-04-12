package top

import "time"

const (
	kGatewayUrl        = "https://gw.api.taobao.com/router/rest"
	kFormat            = "json"
	kApiVersion        = "2.0"
	kPartnerId         = "top-sdk-go-20191025"
	kTimeFormat        = "2006-01-02 15:04:05"
	kKeySign           = "sign"
	kContentTypeForm   = "application/x-www-form-urlencoded"
	kConnectionClose   = "close"
	kHeaderConnection  = "Connection"
	kHeaderContentType = "Content-Type"
	kHeaderUserAgent   = "User-Agent"
	kMethodGet         = "GET"
	kMethodPost        = "POST"
	kHttpClientTimeout = 30 * time.Second
)
