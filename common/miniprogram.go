// 微信小程序
package common

import (
	"fmt"
	"strings"

	"github.com/bigrocs/wechat/responses"
	"github.com/bigrocs/wechat/util"
)

var apiUrlsMiniprogram = map[string]string{
	"auth.code2Session":   "/sns/jscode2session?appid={appid}&secret={secret}&grant_type=authorization_code",
	"auth.getPaidUnionId": "/wxa/getpaidunionid?",
}

// Miniprogram 公共封装
type Miniprogram struct {
	c *Common
}

// ApiUrl 创建 ApiUrl
func (m *Miniprogram) ApiUrl() (apiUrl string, err error) {
	c := m.c.Config
	req := m.c.Requests
	if u, ok := apiUrlsMiniprogram[req.ApiName]; ok {
		u = m.c.APIBaseURL() + u
		apiUrl = strings.Replace(u, "{appid}", c.AppId, -1)
		apiUrl = strings.Replace(apiUrl, "{secret}", c.Secret, -1)
		for key, val := range req.QueryParams {
			apiUrl = apiUrl + "&" + key + "=" + util.InterfaceToString(val)
		}
	} else {
		err = fmt.Errorf("ApiName 不存在请检查。")
	}
	return
}

// Request 执行请求
func (m *Miniprogram) Request(response *responses.CommonResponse) (err error) {
	apiUrl, err := m.ApiUrl()
	if err != nil {
		return err
	}
	res, err := util.HTTPGet(apiUrl)
	if err != nil {
		return err
	}
	response.SetHttpContent(res, "string")
	return
}
