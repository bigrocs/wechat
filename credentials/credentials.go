package credentials

// BaseCredential Deprecated: Use AccessKeyCredential in this package instead.
type BaseCredential struct {
	Miniprogram *Miniprogram
}

// Miniprogram 小程序凭证
type Miniprogram struct {
	AppId  string
	Secret string
}
