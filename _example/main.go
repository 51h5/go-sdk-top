package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	top "github.com/51h5/go-sdk-top"
	"github.com/51h5/go-sdk-top/internal/constants"
	"github.com/51h5/go-sdk-top/request"
	"github.com/51h5/go-sdk-top/response"
)

const (
	appKey = ""
	secret = ""
	//gateway = "https://pre-gw.api.taobao.com/top/router/rest"
)

func main() {
	c := top.New(appKey, secret, 30*time.Second)
	//c.SetGateway(gateway)
	c.SetSignType(constants.SIGN_TYPE_MD5)
	//c.SetSignType(constants.SIGN_TYPE_HMAC_SHA256)

	c.Debug(true)

	// 自定义 http client
	c.SetHttpClient(&http.Client{
		Timeout: 45 * time.Second, // 整体超时 45秒 (连接超时 15秒 + 其他 30秒)
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   15 * time.Second, // 连接超时 15秒
				KeepAlive: 30 * time.Second,
				// DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	})

	grant(c)
	mixnick(c)
	avatar(c)
	tt()
}

func grant(c *top.Client) {
	openid := ""
	bizName := ""
	bizScene := ""
	amount := int64(2)
	actId := ""
	collectionIds := []string{""}
	propertyId := ""

	req := &request.AlibabaAlscGrowthInteractiveMiniGameIntegralGrantRequest{}
	req.MiniGameGrantIntegralRequest = &request.MiniGameGrantIntegralRequest{
		Amount:        amount,
		ActId:         actId,
		BizScene:      bizScene,
		CollectionIds: collectionIds,
		PropertyId:    propertyId,
		RequestId:     fmt.Sprintf("%s-%s-%d-%d-%d", bizScene, openid, time.Now().UnixMilli(), 1, amount),
		OpenId:        openid,
	}
	req.ExtParams = &request.MiniGameGrantIntegralExtParams{BizName: bizName}
	res := &response.AlibabaAlscGrowthInteractiveMiniGameIntegralGrantResponse{}

	code, err := c.Execute(req, res, "")
	if err != nil {
		fmt.Printf("<mixnick> 调用异常: code=%v, err=%s\n", code, err)
		fmt.Printf("reqeust: %v\n", req)
		fmt.Printf("response: %v\n", res)
	}

	if res.Success() {
		fmt.Printf("<grant> 调用成功: %d/%d\n", res.Data.RealGrantValue, res.Data.AccountValue)
	} else {
		fmt.Printf("<grant> 调用失败: topError=%+v\n", res.Error)
		fmt.Printf("<grant> 调用失败: data=%+v\n", res.Data)
	}
}

func mixnick(c *top.Client) {
	req := &request.TaobaoMixnickGetRequest{
		Nick: "放平心态",
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
		Nick: "放01aaNwGavQP79p8UUlM9INu78UWI4RCTeNEy3v+0dzB9E=",
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
	fmt.Printf("nsec(1604848716632) = %s\n", time.Unix(0, 1604848716632*1000000).Format("2006-01-02 15:04:05"))
	fmt.Printf("sec(1604848716632) = %d\n", time.Now().Unix())
	fmt.Printf("sec(1604848716632) = %d\n", time.Now().UnixNano())
	fmt.Printf("Duration(1604848716632) = %v\n", time.Duration(1604848716632))
	fmt.Printf("Milliseconds(1604848716632) = %v\n", time.Duration(1604848716632).Milliseconds())
	fmt.Printf("Milliseconds(1604848716632) = %v\n", (time.Duration(1604848716632) * time.Millisecond).Milliseconds())
	fmt.Printf("Milliseconds(1604848716632) = %v\n", (time.Duration(1604848716632) * time.Millisecond).Nanoseconds())
}
