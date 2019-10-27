package request

import "net/url"

// TOP协议请求 - 共享基类
type topRequest struct {
    targetAppKey string `json:"target_app_key,omitempty"`
}

func (p *topRequest) Check() (code uint, err error) {
    return
}

func (p *topRequest) Method() string {
    return ""
}

func (p *topRequest) Values() url.Values {
    return nil
}

func (p *topRequest) Body() []byte {
    return nil
}

func (p *topRequest) SetBody(v []byte) {}

func (p *topRequest) TargetAppKey() string {
    return p.targetAppKey
}

func (p *topRequest) SetTargetAppKey(v string) {
    p.targetAppKey = v
}
