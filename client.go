package sdk

import (
	"net/http"

	"github.com/bigrocs/wechat/credentials"
	"github.com/bigrocs/wechat/requests"
	"github.com/bigrocs/wechat/responses"
	"github.com/bigrocs/wechat/url"
	"github.com/bigrocs/wechat/util"
)

// Client the type Client
type Client struct {
	httpClient *http.Client
	Credential *credentials.BaseCredential
}

// NewClient 创建默认连接
func NewClient() (client *Client, err error) {
	client = &Client{}
	client.Credential = &credentials.BaseCredential{
		Miniprogram: &credentials.Miniprogram{},
	}
	err = nil
	return
}

// ProcessCommonRequest 处理公共请求
func (client *Client) ProcessCommonRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	response = responses.NewCommonResponse()
	err = client.DoAction(request, response)
	return
}

// DoAction 执行动作
func (client *Client) DoAction(request *requests.CommonRequest, response *responses.CommonResponse) (err error) {
	// 创建访问链接
	u := &url.CommonUrl{
		Credential: client.Credential,
		Requests:   request,
	}
	url, err := u.Url()
	if err != nil {
		return err
	}
	err = client.HTTPGet(url, response)
	if err != nil {
		return err
	}
	return
}

// HTTPGet 请求
func (client *Client) HTTPGet(url string, response *responses.CommonResponse) (err error) {
	res, err := util.HTTPGet(url)
	response.SetHttpContentString(string(res))
	return
}
