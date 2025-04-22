package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type PaymentSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
type WxPayConf struct {
	MchId      string `mapstructure:"MchId"`      //微信商户id
	SerialNo   string `mapstructure:"SerialNo"`   //商户证书的证书序列号
	APIv3Key   string `mapstructure:"APIv3Key"`   //apiV3Key，商户平台获取
	PrivateKey string `mapstructure:"PrivateKey"` //privateKey：私钥 apiclient_key.pem 读取后的内容
	NotifyUrl  string `mapstructure:"NotifyUrl"`  //支付通知回调服务端地址
}

type WxMiniConf struct {
	AppId  string `mapstructure:"AppId"`  //微信小程序appId（非公众号）
	Secret string `mapstructure:"Secret"` //微信小程序secret（非公众号）
}

type ServerConfig struct {
	Name             string           `mapstructure:"name" json:"name"`
	Host             string           `mapstructure:"host" json:"host"`
	OtelInfo         OtelConfig       `mapstructure:"otel" json:"otel"`
	JWTInfo          JWTConfig        `mapstructure:"jwt" json:"jwt"`
	WxMiniConf       WxMiniConf       `mapstructure:"wx_mini" json:"wx_mini"`
	WxPayConf        WxPayConf        `mapstructure:"wx_pay" json:"wx_pay"`
	OrderSrv         OrderSrvConfig   `mapstructure:"order_srv" json:"order_srv"`
	PaymentSrvConfig PaymentSrvConfig `mapstructure:"payment_srv" json:"payment_srv"`
	UserSrv          UserSrvConfig    `mapstructure:"user_srv" json:"user_srv"`
}
