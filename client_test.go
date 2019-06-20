package sdk

import (
	"fmt"
	"testing"

	"github.com/bigrocs/wechat/requests"
)

func TestCode2Session(t *testing.T) {
	// 创建连接
	client, err := NewClient("wx23333bea7abf4245", "efd93e41df66f343b90c42e0ce383eb2")
	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "miniprogram"
	request.ApiName = "auth.code2Session"
	request.QueryParams = map[string]string{
		"js_code": "123456",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req := response.GetHttpContentString()
	fmt.Println(response, err)
	fmt.Println(req, err)
}
