package url

import (
	"github.com/bigrocs/wechat/credentials"
	"github.com/bigrocs/wechat/requests"
)


// CommonUrl 公共封装 
type CommonUrl struct {
	Credential 	*credentials.BaseCredential
	Requests   	*requests.CommonRequest
} 
// Url 创建新的公共连接
func (c *CommonUrl)Url() (url string, err error) {
	req := c.Requests
	// 根据作用域分发
	switch req.Domain {
	case "miniprogram":
		// 小程序
		m := &Miniprogram{c}
		url, err = m.Url()
	}
	return
}
