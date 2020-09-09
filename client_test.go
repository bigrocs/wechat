package wechat

import (
	"fmt"
	"testing"

	"github.com/bigrocs/wechat/requests"
)

func TestCode2Session(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wxa41523f8f3121d3r8f7"
	client.Config.Secret = "4e472dc94711w8b16b891381v6d21422015f"

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

// func TestMchPayMicropay(t *testing.T) {
// 	// 创建连接
// 	client := NewClient()
// 	client.Config.AppId = "wx72ddcf03d5a5fdfe"
// 	client.Config.MchId = "1584521731"
// 	client.Config.ApiKey = "f61ee32da86000a1b16119955fc60eca"
// 	// client.Config.SubAppId = "wx72ddcf03d5a5fdfe"
// 	client.Config.SubMchId = "1597690221"

// 	// 配置参数
// 	request := requests.NewCommonRequest()
// 	request.Domain = "mch"
// 	request.ApiName = "pay.micropay"
// 	request.QueryParams = map[string]interface{}{
// 		"auth_code":        "134550779079187440",
// 		"body":             "测试商品名称1",
// 		"out_trade_no":     "2020021000071",
// 		"total_fee":        "1",
// 		"spbill_create_ip": "127.0.0.1",
// 	}
// 	// 请求
// 	response, err := client.ProcessCommonRequest(request)
// 	req, err := response.GetHttpContentMap()
// 	// fmt.Println(response, err)
// 	// fmt.Println(req, err)
// 	t.Log(req, err)
// }

func TestMchPayRefund(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wx72ddcfgxer5fdfe"
	client.Config.MchId = "15fgxer1731"
	client.Config.ApiKey = "f61ee32da8fgxer119955fc60eca"
	// client.Config.SubAppId = "wx72ddfgxer5a5fdfe"
	client.Config.SubMchId = "159fgxer221"
	client.Config.CA = "/apiclient_cert.p12"
	client.Config.PemCert = ``
	client.Config.PemKey = ``
	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.refund"
	request.QueryParams = map[string]interface{}{
		"out_trade_no":  "2020021000071",
		"out_refund_no": "2020021000071" + "1",
		"total_fee":     "1",
		"refund_fee":    "1",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, e := response.GetHttpContentMap()
	fmt.Println(req, err, e)
}
