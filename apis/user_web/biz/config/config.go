package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type WxMiniConfig struct {
	AppId     string `mapstructure:"app_id" json:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret"`
}
type ServerConfig struct {
	Name       string        `mapstructure:"name" json:"name"`
	Host       string        `mapstructure:"host" json:"host"`
	OtelInfo   OtelConfig    `mapstructure:"otel" json:"otel"`
	UserSrv    UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo    JWTConfig     `mapstructure:"jwt" json:"jwt"`
	WxMiniConf WxMiniConfig  `mapstructure:"wx_mini" json:"wx_mini"`
}
