package wechat

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigrocs/wechat/requests"
)

func TestCode2Session(t *testing.T) {
	// 创建连接
	// client := NewClient()
	// client.Config.AppId =  os.Getenv("PAY_WECHAT_APPID")
	// client.Config.Secret =  os.Getenv("PAY_WECHAT_APPID")
	// client.Config.SignType = "MD5"

	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.Domain = "miniprogram"
	// request.ApiName = "auth.code2Session"
	// request.QueryParams = map[string]interface{}{
	// 	"js_code": "071zClGa1CiEKz0IltHa1KtPeE1zClGx",
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetHttpContentMap()
	// fmt.Println("_____________", req)
	// t.Log(response, err)
	// t.Log(req, err)
}

func TestMchPayMicropay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_WECHAT_APPID")
	client.Config.MchId = os.Getenv("PAY_WECHAT_MCHID")
	client.Config.ApiKey = os.Getenv("PAY_WECHAT_APIKEY")
	// client.Config.SubAppId = os.Getenv("PAY_WECHAT_SUB_APP_ID")
	client.Config.SubMchId = os.Getenv("PAY_WECHAT_SUB_MCHID")

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.micropay"
	request.QueryParams = map[string]interface{}{
		"auth_code":        "136568132728838095",
		"body":             "测试商品名称1",
		"out_trade_no":     "202002100007156",
		"total_fee":        "1",
		"spbill_create_ip": "127.0.0.1",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	// fmt.Println(response, err)
	fmt.Println("TestMchPayMicropay_____", req, err)
	t.Log(req, err)
}

func TestMchPayOrderQuery(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_WECHAT_APPID")
	client.Config.MchId = os.Getenv("PAY_WECHAT_MCHID")
	client.Config.ApiKey = os.Getenv("PAY_WECHAT_APIKEY")
	// client.Config.SubAppId = os.Getenv("PAY_WECHAT_SUB_APP_ID")
	client.Config.SubMchId = os.Getenv("PAY_WECHAT_SUB_MCHID")

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.orderquery"
	request.QueryParams = map[string]interface{}{
		"out_trade_no": "20200210000715",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	// fmt.Println(response, err)
	fmt.Println("TestMchPayOrderQuery_____", req, err)
	t.Log(req, err)
}

func TestMchPayRefund(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_WECHAT_APPID")
	client.Config.MchId = os.Getenv("PAY_WECHAT_MCHID")
	client.Config.ApiKey = os.Getenv("PAY_WECHAT_APIKEY")
	// client.Config.SubAppId = os.Getenv("PAY_WECHAT_SUB_APP_ID")
	client.Config.SubMchId = os.Getenv("PAY_WECHAT_SUB_MCHID")
	// client.Config.CA = "/apiclient_cert.p12"
	client.Config.PemCert = os.Getenv("PAY_WECHAT_PEMCERT")
	client.Config.PemKey = os.Getenv("PAY_WECHAT_PEMKEY")
	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.refund"
	request.QueryParams = map[string]interface{}{
		"out_trade_no":  "202002100007156",
		"out_refund_no": "202002100007156" + "_2",
		"total_fee":     "1",
		"refund_fee":    "1",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, e := response.GetVerifySignDataMap()
	fmt.Println("TestMchPayRefund_____", req, err, e)
}

func TestMchPayRefundQuery(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_WECHAT_APPID")
	client.Config.MchId = os.Getenv("PAY_WECHAT_MCHID")
	client.Config.ApiKey = os.Getenv("PAY_WECHAT_APIKEY")
	// client.Config.SubAppId = os.Getenv("PAY_WECHAT_SUB_APP_ID")
	client.Config.SubMchId = os.Getenv("PAY_WECHAT_SUB_MCHID")
	// client.Config.CA = "/apiclient_cert.p12"
	client.Config.PemCert = os.Getenv("PAY_WECHAT_PEMCERT")
	client.Config.PemKey = os.Getenv("PAY_WECHAT_PEMKEY")
	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.refundquery"
	request.QueryParams = map[string]interface{}{
		"out_trade_no": "2020021000071511A",
		// "out_refund_no": "202002100007151" + "_1",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, e := response.GetVerifySignDataMap()
	fmt.Println("TestMchPayRefundQuery_____", req, err, e)
}

func TestQueryOpenid(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_WECHAT_APPID")
	client.Config.MchId = os.Getenv("PAY_WECHAT_MCHID")
	client.Config.ApiKey = os.Getenv("PAY_WECHAT_APIKEY")
	// client.Config.SubAppId = os.Getenv("PAY_WECHAT_SUB_APP_ID")
	client.Config.SubMchId = os.Getenv("PAY_WECHAT_SUB_MCHID")
	// client.Config.CA = "/apiclient_cert.p12"
	client.Config.PemCert = os.Getenv("PAY_WECHAT_PEMCERT")
	client.Config.PemKey = os.Getenv("PAY_WECHAT_PEMKEY")
	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "tools.authcodetoopenid"
	request.QueryParams = map[string]interface{}{
		"auth_code": "136612636574445977",
		// "out_refund_no": "202002100007151" + "_1",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, e := response.GetVerifySignDataMap()
	fmt.Println("TestQueryOpenid_____", req, err, e)
}

func TestOffiAccountAccessToken(t *testing.T) {
	// 创建连接
	// client := NewClient()
	// client.Config.AppId =  os.Getenv("PAY_WECHAT_APPID")
	// client.Config.Secret =  os.Getenv("PAY_WECHAT_APPID")

	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.Domain = "offiaccount"
	// request.ApiName = "cgi-bin.access_token"
	// // 请求
	// response, e := client.ProcessCommonRequest(request)
	// req, err := response.GetHttpContentMap()
	// fmt.Println("access_token", e, req, err)
	// t.Log(response, err)
	// t.Log(req, err)
}

func TestOffiAccountAuthorize(t *testing.T) {
	// 创建连接
	// 	client := NewClient()
	// 	client.Config.AppId =  os.Getenv("PAY_WECHAT_APPID")
	// 	client.Config.Secret = ""

	// 	// 配置参数
	// 	request := requests.NewCommonRequest()
	// 	request.Domain = "offiaccount"
	// 	request.ApiName = "connect.oauth2.authorize"
	// 	request.QueryParams = map[string]interface{}{
	// 		"redirect_uri": "http://www.xilewanggou.com/rpc",
	// 	}
	// 	// 请求
	// 	response, e := client.ProcessCommonRequest(request)
	// 	req, err := response.GetHttpContentMap()
	// 	fmt.Println("connect.oauth2.authorize", e, req, err)
	// 	t.Log(response, err)
	// 	t.Log(req, err)
}

func TestOffiAccountSnsAccessToken(t *testing.T) {
	// // 创建连接
	// client := NewClient()
	// client.Config.AppId = os.Getenv("PAY_WECHAT_APPID")
	// client.Config.Secret =  os.Getenv("PAY_WECHAT_APPID")

	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.Domain = "offiaccount"
	// request.ApiName = "sns.oauth2.access_token"
	// request.QueryParams = map[string]interface{}{
	// 	"code": "021Tss0w3rOR5V2szy2w3K1oT31Tss0L",
	// }
	// // 请求
	// response, e := client.ProcessCommonRequest(request)
	// req, err := response.GetHttpContentMap()
	// fmt.Println("sns.oauth2.access_token", e, req, err)
	// t.Log(response, err)
	// t.Log(req, err)
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
