package sdk

import (
	"fmt"
	"net/http"

	"github.com/bigrocs/wechat/credentials"
	"github.com/bigrocs/wechat/requests"
	"github.com/bigrocs/wechat/responses"
)

// Client the type Client
type Client struct {
	httpClient *http.Client
	credential *credentials.BaseCredential
}

// NewClient 创建默认连接
func NewClient(appId, secret string) (client *Client, err error) {
	client = &Client{}
	client.credential = &credentials.BaseCredential{
		AppId:  appId,
		Secret: secret,
	}
	err = nil
	return
}

// ProcessCommonRequest 处理公共请求
func (client *Client) ProcessCommonRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	fmt.Println("aa")
	response = responses.NewCommonResponse()
	fmt.Println(client.credential, request)
	return
}
