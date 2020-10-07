package wechat

import (
	"fmt"
	"testing"

	"github.com/bigrocs/wechat/requests"
)

func TestCode2Session(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wx15550c1a89d982c8"
	client.Config.Secret = "f9c11f183a5beb592ccd801298ff5533"

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "miniprogram"
	request.ApiName = "auth.code2Session"
	request.QueryParams = map[string]interface{}{
		"js_code": "011e54Ga1XzcIz027NGa1ZfYOF2e54GO",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	fmt.Println("_____________", req)
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
	// client := NewClient()
	// client.Config.AppId = "wx72ddcfgxer5fdfe"
	// client.Config.MchId = "15fgxer1731"
	// client.Config.ApiKey = "f61ee32da8fgxer119955fc60eca"
	// // client.Config.SubAppId = "wx72ddfgxer5a5fdfe"
	// client.Config.SubMchId = "159fgxer221"
	// // client.Config.CA = "/apiclient_cert.p12"
	// client.Config.PemCert = ``
	// client.Config.PemKey = ``
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.Domain = "mch"
	// request.ApiName = "pay.refund"
	// request.QueryParams = map[string]interface{}{
	// 	"out_trade_no":  "2020021000071",
	// 	"out_refund_no": "2020021000071" + "1",
	// 	"total_fee":     "1",
	// 	"refund_fee":    "1",
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, e := response.GetHttpContentMap()
	// fmt.Println(req, err, e)
}

func TestOffiAccountAccessToken(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wx15550c1a89d982c8"
	client.Config.Secret = "f9c11f183a5beb592ccd801298ff5533"

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "offiaccount"
	request.ApiName = "cgi-bin.access_token"
	// 请求
	response, e := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	fmt.Println("access_token", e, req, err)
	t.Log(response, err)
	t.Log(req, err)
}

func TestOffiAccountMessageTemplate(t *testing.T) {
	// // 创建连接
	// client := NewClient()
	// client.Config.AccessToken = "37_rZB3K_VGcf-4z--ppMEj9KtAzKKjucdwAqT7ylA0XVGeAHAQOZ"
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.Domain = "offiaccount"
	// request.ApiName = "message.template"
	// request.QueryParams = map[string]interface{}{
	// 	"touser":      "",
	// 	"template_id": "ybgOF-ZQsWTr8JS0lGwuRzFPdBKGAsiJiIk5ZX0EaDY",
	// 	"url":         "http://www.xilewanggou.com/download",
	// 	"data": map[string]interface{}{
	// 		"first": map[string]interface{}{
	// 			"value": "恭喜你下单成功！",
	// 			"color": "#173177",
	// 		},
	// 		"keyword1": map[string]interface{}{
	// 			"value": "2020年09月29日 21:49",
	// 			"color": "#173177",
	// 		},
	// 		"keyword2": map[string]interface{}{
	// 			"value": "苹果",
	// 			"color": "#173177",
	// 		},
	// 		"keyword3": map[string]interface{}{
	// 			"value": "00122009280001",
	// 			"color": "#173177",
	// 		},
	// 		"remark": map[string]interface{}{
	// 			"value": "欢迎再次购买！",
	// 			"color": "#173177",
	// 		},
	// 	},
	// }
	// // 请求
	// response, e := client.ProcessCommonRequest(request)
	// req, err := response.GetHttpContentMap()
	// fmt.Println("message.template", e, req, err)
	// t.Log(response, err)
	// t.Log(req, err)
}
