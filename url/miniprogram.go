package url
import (
	"fmt"
	"strings"
)
var urls = map[string]string{
	"auth.code2Session" : "https://api.weixin.qq.com/sns/jscode2session?appid={appid}&secret={secret}&js_code={js_code}&grant_type=authorization_code",
	"auth.getPaidUnionId" : "https://api.weixin.qq.com/wxa/getpaidunionid?access_token={access_token}&openid={openid}&transaction_id={transaction_id}&mch_id={mch_id}&out_trade_no={out_trade_no}",
}


// Miniprogram 公共封装 
type Miniprogram struct {
	c *CommonUrl
}

// Url 创建新的公共连接
func (m *Miniprogram)Url() (url string, err error) {
	c := m.c.Credential
	req := m.c.Requests
	if url, ok := urls[req.ApiName]; ok {
		url = strings.Replace(url, "{appid}", c.AppId, -1)
		url = strings.Replace(url, "{secret}", c.Secret, -1)
		url = strings.Replace(url, "{mch_id}", c.MchId, -1)
		for key, val := range req.QueryParams {
			// 替换全部参数 {}
			url = strings.Replace(url, "{"+key+"}", val, -1)
		}
	}else{
		err = fmt.Errorf("ApiName 不存在,请检查。")
	}
	return
}