package hpingan

type Config struct {
	BaseServiceUrl    string `json:"baseServiceUrl" description:"平安银行接口服务URL"`
	Channel           string `json:"channel" description:"平安分配的渠道编号"`
	ChannelPrivateKey string `json:"channelPrivateKey" description:"渠道RSA私钥Base64值，平安银行提供"`
}
