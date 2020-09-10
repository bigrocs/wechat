# wechat SDK
### 微服务内核访问演示
#### 示例
##### main.go
```
package main

import (
	"github.com/bigrocs/wechat"
    "github.com/bigrocs/wechat/requests"
)
func main() {
	client := wechat.NewClient()
	client.Config.AppId = ""
	client.Config.MchId = ""
	client.Config.ApiKey = ""
	// client.Config.SubAppId = ""
	client.Config.SubMchId = ""
	client.Config.CA = "/apiclient_cert.p12"
	client.Config.PemCert = ``
	client.Config.PemKey = ``
    // 退款示例 需要ca证书或pem证书
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
```
#### 小程序登陆授权
```
{
    "Domain": "miniprogram",
    "ApiName": "auth.code2Session",
    "QueryParams": {
    	"js_code":"0211MY2j255VoB0WKO3j2cqJ2j21MY2s"
    }
}
```
#### 微信付款码支付
```
{
    "Domain": "mch",
    "ApiName": "pay.micropay",
    "QueryParams": {
    	"auth_code":        "134770030978364234",
        "body":             "测试商品名称1",
		"out_trade_no":     "202002100007",
		"total_fee":        "1",
		"spbill_create_ip": "127.0.0.1",
    }
}
```
#### 微信付款码支付查询
```
{
    "Domain": "mch",
    "ApiName": "pay.orderquery",
    "QueryParams": {
		"out_trade_no":     "202002100007",
    }
}
```
- 具体参数参考微信开发文档
- https://api.wechat.com/
