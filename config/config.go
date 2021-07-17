package config

type Config struct {
	AppId       string // 公众号AppId 服务商的APPID
	MchId       string // 商户号
	ApiKey      string // key为商户平台设置的密钥key
	SignType    string // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	SubAppId    string // 子商户公众账号AppId
	SubMchId    string // 子商户号
	CA          string // 双向证书
	PemCert     string // pem 双向证书
	PemKey      string // pem 双向证书
	Secret      string
	AccessToken string // 网页版授权Token
	SessionKey  string // 小程序会话密钥
}
