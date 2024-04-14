package request

import (
	"net/url"

	"github.com/51h5/go-sdk-top/internal/utils"
	jsoniter "github.com/json-iterator/go"
)

type AlibabaAlscGrowthInteractiveMiniGameIntegralGrantRequest struct {
	topRequest

	// 小游戏积分发放请求参数
	MiniGameGrantIntegralRequest *MiniGameGrantIntegralRequest `json:"mini_game_grant_integral_request"`

	// 饿了么项目, 基本只需传 `BizName` 即可
	ExtParams *MiniGameGrantIntegralExtParams `json:"-,omitempty"`
}

// 小游戏积分发放请求参数
type MiniGameGrantIntegralRequest struct {
	Amount        int64    `json:"amount"`                // 积分发放数量
	ActId         string   `json:"act_id"`                // 活动id
	BizScene      string   `json:"biz_scene"`             // 业务场景
	CollectionIds []string `json:"collection_ids"`        // 组建集id集合
	PropertyId    string   `json:"property_id"`           // 资产id
	RequestId     string   `json:"request_id"`            // 请求id(幂等id，非链路id),不超过128位,相同幂等号,发放数量也要保持相同，否则会幂等异常
	OpenId        string   `json:"open_id,omitempty"`     // openId，非成语场景下必填
	GameAccId     string   `json:"game_acc_id,omitempty"` // 游戏账号，成语场景下必填
	ExtParams     string   `json:"ext_params,omitempty"`  // 按实际情况填写相关参数(非业务参数,从奇门dispatch接口转发过去的最早h5方式,可以直接把透传过去的参数再透传过来,小程序方式接入的从header中获取，透传参数具体参考下述链接对应的文档：https://bdshhd.yuque.com/org-wiki-bdshhd-refgtt/white-paper/zwf4r3v9hxhaabxa
}

// 小游戏积分发放请求参数 的 扩展参数
type MiniGameGrantIntegralExtParams struct {
	BizName         string `json:"bizName"`              // 游戏名称用于乐园币明细展示,中文游戏名称
	GameStage       string `json:"gameStage,omitempty"`  // 什么场景下发放的乐园币(具体参考描述中链接的文档)
	GameOperateName string `json:"gameStage,omitempty"`  // 如果是 "任务展示"、"完成任务领取奖励"环节返回任务名称
	GameTaskId      string `json:"gameTaskId,omitempty"` // 和任务相关的，给任务id
	Ua              string `json:"ua,omitempty"`         // h5从奇门dispatch透传参数中取,小程序从header取
	Wua             string `json:"wua,omitempty"`        // h5从奇门dispatch透传参数中取,小程序从header取
	XMiniUa         string `json:"x-mini-ua,omitempty"`  // h5从奇门dispatch透传参数中取,小程序从header取
	XUmt            string `json:"x-umt,omitempty"`      // h5从奇门dispatch透传参数中取,小程序从header取
}

func (r *AlibabaAlscGrowthInteractiveMiniGameIntegralGrantRequest) Method() string {
	return "alibaba.alsc.growth.interactive.mini.game.integral.grant"
}

func (r *AlibabaAlscGrowthInteractiveMiniGameIntegralGrantRequest) Values() url.Values {
	if r.ExtParams != nil {
		ext, _ := jsoniter.Marshal(r.ExtParams)
		r.MiniGameGrantIntegralRequest.ExtParams = string(ext)
	}
	bs, _ := jsoniter.Marshal(r.MiniGameGrantIntegralRequest)

	return url.Values{
		"mini_game_grant_integral_request": {string(bs)},
	}
}

func (r *AlibabaAlscGrowthInteractiveMiniGameIntegralGrantRequest) Check() (code uint, err error) {
	code, err = utils.CheckNotEmpty("mini_game_grant_integral_request.act_id", r.MiniGameGrantIntegralRequest.ActId)
	if err != nil {
		return
	}

	code, err = utils.CheckNotEmpty("mini_game_grant_integral_request.biz_scene", r.MiniGameGrantIntegralRequest.BizScene)
	if err != nil {
		return
	}

	code, err = utils.CheckNotEmpty("mini_game_grant_integral_request.property_id", r.MiniGameGrantIntegralRequest.PropertyId)
	if err != nil {
		return
	}

	code, err = utils.CheckNotEmpty("mini_game_grant_integral_request.request_id", r.MiniGameGrantIntegralRequest.RequestId)
	if err != nil {
		return
	}

	return
}
