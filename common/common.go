package common

import (
	"github.com/bigrocs/wechat/config"
	"github.com/bigrocs/wechat/requests"
	"github.com/bigrocs/wechat/responses"
)

// Common 公共封装
type Common struct {
	Config   *config.Config
	Requests *requests.CommonRequest
}

// Action 创建新的公共连接
func (c *Common) Action(response *responses.CommonResponse) (err error) {
	req := c.Requests
	// 根据作用域分发
	switch req.Domain {
	case "offiaccount": // 小程序
		h := &OffiAccount{c}
		err = h.Request(response)
	case "miniprogram": // 小程序
		h := &Miniprogram{c}
		err = h.Request(response)
	case "mch": // 微信商户平台(微信支付) SDK
		// 小程序
		h := &Mch{c}
		err = h.Request(response)
	}
	return
}

// 默认 API
func (c *Common) APIBaseURL() string { // TODO(): 后期做容灾功能
	return "https://api.weixin.qq.com"
}

// 公众号 API
func (c *Common) APIBaseURLOffiAccount() string { // TODO(): 后期做容灾功能
	return "https://api.weixin.qq.com"
}

// // 微信商户平台(微信支付) API
func (c *Common) APIBaseURLMch() string { // TODO(): 后期做容灾功能
	return "https://api.mch.weixin.qq.com"
}
