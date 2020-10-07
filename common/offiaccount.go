// 微信小g公众号
package common

import (
	"fmt"
	"strings"

	"github.com/bigrocs/wechat/responses"
	"github.com/bigrocs/wechat/util"
)

// apiUrl 请求api结构
type apiUrl struct {
	Name   string
	Method string
	URL    string
}

var apiUrlsOffiAccount = []apiUrl{
	{ // 通过code换取网页授权access_token
		Name:   "sns.oauth2.access_token",
		Method: "get",
		URL:    "/sns/oauth2/access_token?appid={appid}&secret={secret}&grant_type=authorization_code",
	}, { // 刷新access_token（如果需要）
		Name:   "sns.oauth2.refresh_token",
		Method: "get",
		URL:    "/sns/oauth2/refresh_token?appid={appid}&grant_type=refresh_token",
	}, { // 拉取用户信息(需scope为 snsapi_userinfo)
		Name:   "sns.userinfo",
		Method: "get",
		URL:    "/sns/userinfo?lang=zh_CN",
	}, { // 检验授权凭证（access_token）是否有效
		Name:   "sns.auth",
		Method: "get",
		URL:    "/sns/auth?",
	},
	{
		Name:   "cgi-bin.access_token",
		Method: "get",
		URL:    "/cgi-bin/token?grant_type=client_credential&appid={appid}&secret={secret}",
	}, {
		Name:   "cgi-bin.message.template.send",
		Method: "post",
		URL:    "/cgi-bin/message/template/send?access_token={access_token}",
	},
}

// OffiAccount 公共封装
type OffiAccount struct {
	c *Common
}

// ApiUrl 创建 ApiUrl
func (o *OffiAccount) ApiUrl() (apiUrl string, method string, err error) {
	c := o.c.Config
	req := o.c.Requests
	for _, u := range apiUrlsOffiAccount {
		if u.Name == req.ApiName {
			method = u.Method
			url := o.c.APIBaseURLOffiAccount() + u.URL
			apiUrl = strings.Replace(url, "{appid}", c.AppId, -1)
			apiUrl = strings.Replace(apiUrl, "{secret}", c.Secret, -1)
			apiUrl = strings.Replace(apiUrl, "{access_token}", c.AccessToken, -1)
			if u.Method == "get" {
				for key, val := range req.QueryParams {
					apiUrl = apiUrl + "&" + key + "=" + util.InterfaceToString(val)
				}
			}
		}
	}
	if apiUrl == "" {
		err = fmt.Errorf("ApiName 不存在请检查。")
	}
	return
}

// Request 执行请求
func (o *OffiAccount) Request(response *responses.CommonResponse) (err error) {
	apiUrl, method, err := o.ApiUrl()
	if err != nil {
		return err
	}
	req := o.c.Requests
	var res []byte
	if method == "post" {
		res, err = util.PostJSON(apiUrl, req.QueryParams)
		if err != nil {
			return err
		}
	}
	if method == "get" {
		res, err = util.HTTPGet(apiUrl)
		if err != nil {
			return err
		}
	}
	response.SetHttpContent(res, "string")
	return
}
