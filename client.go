package top

import (
    "bytes"
    "fmt"
    jsoniter "github.com/json-iterator/go"
    "gitlab.51h5.com/go/sdk-top/internal/constants"
    "gitlab.51h5.com/go/sdk-top/internal/utils"
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
)

type Client struct {
    debug        bool
    checkRequest bool
    keepAlive    bool
    ua           string
    appKey       string
    secretKey    string
    signType     string
    gateway      string
    httpClient   *http.Client
}

func New(appKey, secretKey string, timeout time.Duration) (client *Client) {
    client = &Client{
        // debug:      false,
        keepAlive:    true,
        checkRequest: true,
        appKey:       appKey,
        secretKey:    secretKey,
        gateway:      kGatewayUrl,
        signType:     constants.SIGN_TYPE_MD5,
    }

    if timeout > 0 {
        client.httpClient = &http.Client{
            Timeout: timeout,
        }
    } else {
        client.httpClient = &http.Client{
            Timeout: kHttpClientTimeout,
        }
    }

    return
}

// 调用接口
// - session: 用户的授权令牌
// - target_app_key: 被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效
func (c *Client) Execute(req Request, res Response, session string) (code uint, err error) {
    if c.checkRequest {
        code, err = req.Check()
        if err != nil {
            return
        }
    }

    // 公共参数
    sv := url.Values{
        "app_key":     {c.appKey},
        "format":      {kFormat},
        "sign_method": {c.signType},
        "v":           {kApiVersion},
        "timestamp":   {time.Now().Format(kTimeFormat)},
        "partner_id":  {kPartnerId},
        // "simplify":       {"true"},
        // "session":        {session},
        // "target_app_key": {req.TargetAppKey()},
        "method": {req.Method()},
    }

    if session != "" {
        sv.Set("session", session)
    }

    if req.TargetAppKey() != "" {
        sv.Set("target_app_key", req.TargetAppKey())
    }

    if c.debug {
        fmt.Printf("[top.Execute] 公共参数: %s\n", sv.Encode())
    }

    // 业务参数
    av := req.Values()
    if c.debug {
        fmt.Printf("[top.Execute] 业务参数: %s\n", av.Encode())
    }

    // 计算签名
    sv.Set(kKeySign, c.sign(sv, av, req))

    if c.debug {
        fmt.Printf("[top.Execute] 请求地址: %s?%s\n", c.gateway, sv.Encode())
    }

    var r *http.Request
    if req.Body() == nil {
        r, err = http.NewRequest(kMethodGet, c.gateway+"?"+sv.Encode(), nil)
    } else {
        r, err = http.NewRequest(kMethodPost, c.gateway+"?"+sv.Encode(), bytes.NewReader(req.Body()))
        if err == nil {
            r.Header.Set(kHeaderContentType, kContentTypeForm)
        }
    }

    if err != nil {
        code = 997
        return
    }

    if !c.keepAlive {
        r.Header.Set(kHeaderConnection, kConnectionClose)
    }

    if c.ua != "" {
        r.Header.Set(kHeaderUserAgent, c.ua)
    }

    body, err := doRequest(c.httpClient, r)
    if err != nil {
        code = 998
        return
    }

    if c.debug {
        fmt.Printf("[top.Execute] 请求返回: %s\n", string(body))
    }

    err = parseJsonResponse(body, res)
    if err != nil {
        code = 999
        return
    }

    // XXX: 淘宝TOP接口返回结构：混乱 & 垃圾
    res.Fix()

    return
}

func (c *Client) Debug(enable bool) {
    c.debug = enable
}

func (c *Client) CheckRequest(enable bool) {
    c.checkRequest = enable
}

func (c *Client) SetKeepAlive(v bool) {
    c.keepAlive = v
}

func (c *Client) SetUserAgent(ua string) {
    if ua != "" {
        c.ua = ua
    }
}

func (c *Client) SetGateway(v string) {
    c.gateway = v
}

func (c *Client) SetSignType(v string) {
    c.signType = v
}

func (c *Client) SetHttpClient(hc *http.Client) {
    c.httpClient = hc
}

func (c *Client) sign(sysParams, apiParams url.Values, req Request) string {
    if apiParams != nil {
        for k := range apiParams {
            if apiParams.Get(k) != "" {
                sysParams.Set(k, apiParams.Get(k))
            }
        }
    }

    if c.debug {
        fmt.Printf("[top.sign] 签名串: %v\n", sysParams)
    }

    s := utils.SignToRequest(sysParams, req.Body(), c.secretKey, c.signType)

    if c.debug {
        fmt.Printf("[top.sign] 签名: %s\n", s)
    }

    return s
}

func parseJsonResponse(body []byte, res Response) error {
    return jsoniter.Unmarshal(body, res)
}

func doRequest(c *http.Client, r *http.Request) ([]byte, error) {
    res, err := c.Do(r)
    if res != nil {
        defer res.Body.Close()
    }

    if err != nil {
        return nil, err
    }

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("response code %d", res.StatusCode)
    }

    bits, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    return bits, nil
}
