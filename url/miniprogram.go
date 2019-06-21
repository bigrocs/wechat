package url

import (
	"fmt"
	"strings"
)

var urls = map[string]string{
	"auth.code2Session":   "https://api.weixin.qq.com/sns/jscode2session?appid={appid}&secret={secret}&grant_type=authorization_code",
	"auth.getPaidUnionId": "https://api.weixin.qq.com/wxa/getpaidunionid?",
}

// Miniprogram 公共封装
type Miniprogram struct {
	c *CommonUrl
}

// Url 创建新的公共连接
func (m *Miniprogram) Url() (url string, err error) {
	c := m.c.Credential.Miniprogram
	req := m.c.Requests
	if u, ok := urls[req.ApiName]; ok {
		url = strings.Replace(u, "{appid}", c.AppId, -1)
		url = strings.Replace(url, "{secret}", c.Secret, -1)
		for key, val := range req.QueryParams {
			url = url + "&" + key + "=" + val
		}
	} else {
		err = fmt.Errorf("ApiName 不存在请检查。")
	}
	return
}
