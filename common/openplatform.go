// 微信公众号
package common

import (
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
	{
		// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
		Name:   "api_component_token",
		Method: "post",
		URL:    "https://api.weixin.qq.com/cgi-bin/component/api_component_token",
	},
	{
		// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
		Name:   "api_create_preauthcode",
		Method: "post",
		URL:    "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token={component_access_token}",
	},
	{
		// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html
		Name:   "api_query_auth",
		Method: "post",
		URL:    "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token={component_access_token}",
	},
	{
		// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
		Name:   "api_authorizer_token",
		Method: "post",
		URL:    "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token={component_access_token}",
	},
	{
		// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_get_authorizer_info.html
		Name:   "api_get_authorizer_info",
		Method: "post",
		URL:    "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token={component_access_token}",
	},
	{
		Name:   "componentloginpage",
		Method: "url",
		URL:    "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid={component_appid}&pre_auth_code={pre_auth_code}&redirect_uri={redirect_uri}&auth_type={auth_type}&biz_appid={biz_appid}",
	},
}

// Openplatform 公共封装
type Openplatform struct {
	c *Common
}

// ApiUrl 创建 ApiUrl
func (o *Openplatform) ApiUrl() (apiUrl string, method string, err error) {
	c := o.c.Config
	req := o.c.Requests
	for _, u := range apiUrlsOffiAccount {
		if u.Name == req.ApiName {
			method = u.Method
			apiUrl = u.URL
			apiUrl = strings.Replace(apiUrl, "{appid}", c.AppId, -1)
			apiUrl = strings.Replace(apiUrl, "{secret}", c.Secret, -1)
			// apiUrl = strings.Replace(apiUrl, "{component_appid}", c.AppId, -1)
			// apiUrl = strings.Replace(apiUrl, "{component_appsecret}", c.Secret, -1)
		}
	}
	return
}

// Request 执行请求
func (o *Openplatform) Request(response *responses.CommonResponse) (err error) {
	req := o.c.Requests
	apiUrl, method, err := o.ApiUrl()
	if err != nil {
		return err
	}
	var res []byte
	if method == "post" {
		res, err = util.PostJSON(apiUrl, req.QueryParams)
		if err != nil {
			return err
		}
		response.SetHttpContent(res, "string")
	}
	if method == "get" {
		res, err = util.HTTPGet(apiUrl)
		if err != nil {
			return err
		}
		response.SetHttpContent(res, "string")
	}
	if method == "url" {
		apiUrl = strings.Replace(apiUrl, "{pre_auth_code}", req.QueryParams["pre_auth_code"], -1)
		apiUrl = strings.Replace(apiUrl, "{redirect_uri}", req.QueryParams["redirect_uri"], -1)
		apiUrl = strings.Replace(apiUrl, "{auth_type}", req.QueryParams["auth_type"], -1)
		apiUrl = strings.Replace(apiUrl, "{biz_appid}", req.QueryParams["biz_appid"], -1)
		response.SetHttpContent([]byte(`
		{
			"url": "`+apiUrl+`"
		}
		`), "string")
	}
	return
}
