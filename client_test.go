package wechat

import (
	"fmt"
	"testing"

	"github.com/bigrocs/wechat/requests"
)

func TestCode2Session(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wxa4153f8f312d3r8f7"
	client.Config.Secret = "4e47dc9471w8b16b891381v6d1422015f"

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "miniprogram"
	request.ApiName = "auth.code2Session"
	request.QueryParams = map[string]interface{}{
		"js_code": "0211MY2j255Vo1B0WKO3j2cqJ2j21MY2s",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	t.Log(response, err)
	t.Log(req, err)
}

func TestMchPayMicropay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wxa4153f8f32d13r8f7"
	client.Config.MchId = "14195241271"
	client.Config.ApiKey = "4e47dc9471w8b6b8191381v6d1422015f"
	client.Config.SubAppId = "wx48dc842f51284028c"
	client.Config.SubMchId = "14364311421"

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.micropay"
	request.QueryParams = map[string]interface{}{
		"auth_code":        "134770030978364234",
		"body":             "测试商品名称1",
		"out_trade_no":     "202002100007",
		"total_fee":        "1",
		"spbill_create_ip": "127.0.0.1",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	// fmt.Println(response, err)
	fmt.Println(req, err)
}
