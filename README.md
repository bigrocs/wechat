# wechat SDK
### 微服务内核访问演示

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
- 具体参数参考微信开发文档
- https://api.wechat.com/
