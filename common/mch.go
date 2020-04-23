// 微信商户平台(微信支付) SDK
package common

import (
	"fmt"

	"github.com/bigrocs/wechat/responses"
	"github.com/bigrocs/wechat/util"
)

var apiUrlsMch = map[string]string{
	"pay.micropay":   "/pay/micropay",   //付款码支付
	"pay.orderquery": "/pay/orderquery", //付款码支付查询
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
	c := m.c.Config
	req := m.c.Requests
	apiUrl, err := m.ApiUrl()
	// 构建配置参数
	if _, ok := req.QueryParams["appid"]; !ok {
		req.QueryParams["appid"] = c.AppId
	}
	if _, ok := req.QueryParams["mch_id"]; !ok {
		req.QueryParams["mch_id"] = c.MchId
	}
	if _, ok := req.QueryParams["sub_appid"]; !ok {
		req.QueryParams["sub_appid"] = c.SubAppId
	}
	if _, ok := req.QueryParams["sub_mch_id"]; !ok {
		req.QueryParams["sub_mch_id"] = c.SubMchId
	}

	req.QueryParams["nonce_str"] = util.NonceStr()
	req.QueryParams["sign"] = util.Sign(req.QueryParams, c.ApiKey, util.SignType_MD5)
	if err != nil {
		return err
	}
	res, err := util.PostXML(apiUrl, req.QueryParams)
	if err != nil {
		return err
	}
	response.SetHttpContent(res, "xml")
	return
}
