package config

type Config struct {
	AppId    string // 公众号AppId 服务商的APPID
	MchId    string // 商户号
	ApiKey   string // key为商户平台设置的密钥key
	SubAppId string // 子商户公众账号AppId
	SubMchId string // 子商户号
	Secret   string
	CA       string // 双向证书
}
