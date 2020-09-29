// 微信小g公众号
package common

import (
	"fmt"
	"strings"

	"github.com/bigrocs/wechat/responses"
	"github.com/bigrocs/wechat/util"
)

var apiUrlsOffiAccount = map[string]string{
	"access_token":     "/token?grant_type=client_credential&appid={appid}&secret={secret}",
	"message.template": "/message/template/send?access_token={access_token}",
}

// OffiAccount 公共封装
type OffiAccount struct {
	c *Common
}

// ApiUrl 创建 ApiUrl
func (o *OffiAccount) ApiUrl() (apiUrl string, err error) {
	c := o.c.Config
	req := o.c.Requests
	if u, ok := apiUrlsOffiAccount[req.ApiName]; ok {
		u = o.c.APIBaseURLOffiAccount() + u
		apiUrl = strings.Replace(u, "{appid}", c.AppId, -1)
		apiUrl = strings.Replace(apiUrl, "{secret}", c.Secret, -1)
		apiUrl = strings.Replace(apiUrl, "{access_token}", c.AccessToken, -1)
	} else {
		err = fmt.Errorf("ApiName 不存在请检查。")
	}
	return
}

// Request 执行请求
func (o *OffiAccount) Request(response *responses.CommonResponse) (err error) {
	apiUrl, err := o.ApiUrl()
	req := o.c.Requests
	if err != nil {
		return err
	}
	res, err := util.PostJSON(apiUrl, req.QueryParams)
	if err != nil {
		return err
	}
	response.SetHttpContent(res, "string")
	return
}
