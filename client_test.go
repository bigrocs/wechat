package sdk

import (
	"fmt"
	"testing"

	"github.com/bigrocs/wechat/requests"
)

func TestCode2Session(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = ""
	client.Config.Secret = ""

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "miniprogram"
	request.ApiName = "auth.code2Session"
	request.QueryParams = map[string]string{
		"js_code": "0211MY2j255VoB0WKO3j2cqJ2j21MY2s",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req := response.GetHttpContentString()
	t.Log(response, err)
	t.Log(req, err)
}

func TestMchPayMicropay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = ""
	client.Config.Secret = ""

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.micropay"
	request.QueryParams = map[string]string{
		"js_code": "0211MY2j255VoB0WKO3j2cqJ2j21MY2s",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req := response.GetHttpContentString()
	fmt.Println(response, err)
	fmt.Println(req, err)
}
