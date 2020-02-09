// 微信商户平台(微信支付) SDK
package common

import (
	"fmt"

	"github.com/bigrocs/wechat/responses"
	"github.com/bigrocs/wechat/util"
)

var apiUrlsMch = map[string]string{
	"pay.micropay": "/pay/micropay", //付款码支付
}

// Mch 公共封装
type Mch struct {
	c *Common
}

// ApiUrl 创建 ApiUrl
func (m *Mch) ApiUrl() (apiUrl string, err error) {
	req := m.c.Requests
	if u, ok := apiUrlsMch[req.ApiName]; ok {
		apiUrl = m.c.APIBaseURLMch() + u
	} else {
		err = fmt.Errorf("ApiName 不存在请检查。")
	}
	return
}

// Request 执行请求
func (m *Mch) Request(response *responses.CommonResponse) (err error) {
	req := m.c.Requests
	apiUrl, err := m.ApiUrl()
	queryParams := req.QueryParams
	if err != nil {
		return err
	}
	res, err := util.PostXML(apiUrl, queryParams)
	if err != nil {
		return err
	}
	response.SetHttpContentString(string(res))
	return
}
