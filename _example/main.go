package main

import (
    "51h5.com/sdk/top"
    "51h5.com/sdk/top/request"
    "51h5.com/sdk/top/response"
    "fmt"
    "time"
)

const (
    appKey  = "23638943"
    secret  = ""
    gateway = "https://proxy.hz.taeapp.com/top/"
)

func main() {
    // c := top.New(appKey, secret, 30 * time.Second)
    // c.SetGateway(gateway)
    // c.SetSignType(constants.SIGN_TYPE_MD5)
    //
    // c.Debug(false)

    // // 自定义 http client
    // c.SetHttpClient(&http.Client{
    //     Timeout: 45 * time.Millisecond, // 整体超时 45秒 (连接超时 15秒 + 其他 30秒)
    //     Transport: &http.Transport{
    //         Proxy: http.ProxyFromEnvironment,
    //         DialContext: (&net.Dialer{
    //             Timeout:   15 * time.Second, // 连接超时 15秒
    //             KeepAlive: 30 * time.Second,
    //             // DualStack: true,
    //         }).DialContext,
    //         ForceAttemptHTTP2:     true,
    //         MaxIdleConns:          100,
    //         IdleConnTimeout:       90 * time.Second,
    //         TLSHandshakeTimeout:   10 * time.Second,
    //         ExpectContinueTimeout: 1 * time.Second,
    //     },
    // })

    // mixnick(c)
    // avatar(c)
    tt()
}

func mixnick(c *top.Client) {
    req := &request.TaobaoMixnickGetRequest{
        Nick: "想飞的鱼",
    }
    res := &response.TaobaoMixnickGetResponse{}

    code, err := c.Execute(req, res, "")
    if err != nil {
        fmt.Printf("<mixnick> 调用异常: code=%v, err=%s\n", code, err)
        fmt.Printf("reqeust: %v\n", req)
        fmt.Printf("response: %v\n", res)
    }

    if res.Success() {
        fmt.Printf("<mixnick> 调用成功: %s\n", res.Data.Nick)
    } else {
        fmt.Printf("<mixnick> 调用失败: code=%v, msg=%s, subCode=%s, subMsg=%s\n", res.Error.Code, res.Error.Msg, res.Error.SubCode, res.Error.SubMsg)
    }
}

func avatar(c *top.Client) {
    req := &request.TaobaoUserAvatarGetRequest{
        Nick: "反01FlPrQ9+keP3qZ3IhczpZFByAbDrwQEY7WNLYOaFROMA=",
    }
    res := &response.TaobaoUserAvatarGetResponse{}

    code, err := c.Execute(req, res, "")
    if err != nil {
        fmt.Printf("<avatar> 调用异常: code=%v, err=%s\n", code, err)
        fmt.Printf("reqeust: %v\n", req)
        fmt.Printf("response: %v\n", res)
    }

    if res.Success() {
        fmt.Printf("<avatar> 调用成功: %s\n", res.Data.Avatar)
    } else {
        fmt.Printf("<avatar> 调用失败: code=%v, msg=%s, subCode=%s, subMsg=%s\n", res.Error.Code, res.Error.Msg, res.Error.SubCode, res.Error.SubMsg)
    }
}

func tt() {
    fmt.Printf("sec(1604848716632) = %s\n", time.Unix(1604848716632, 0).Format("2006-01-02 15:04:05"))
    fmt.Printf("nsec(1604848716632) = %s\n", time.Unix(0, 1604848716632).Format("2006-01-02 15:04:05"))
    fmt.Printf("nsec(1604848716632) = %s\n", time.Unix(0, 1604848716632 * 1000000).Format("2006-01-02 15:04:05"))
    fmt.Printf("sec(1604848716632) = %d\n", time.Now().Unix())
    fmt.Printf("sec(1604848716632) = %d\n", time.Now().UnixNano())
    fmt.Printf("Duration(1604848716632) = %v\n", time.Duration(1604848716632))
    fmt.Printf("Milliseconds(1604848716632) = %v\n", time.Duration(1604848716632).Milliseconds())
    fmt.Printf("Milliseconds(1604848716632) = %v\n", (time.Duration(1604848716632) * time.Millisecond).Milliseconds())
    fmt.Printf("Milliseconds(1604848716632) = %v\n", (time.Duration(1604848716632) * time.Millisecond).Nanoseconds())
}