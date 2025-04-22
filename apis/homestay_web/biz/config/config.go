package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type HomestaySrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
type ServerConfig struct {
	Name        string            `mapstructure:"name" json:"name"`
	Host        string            `mapstructure:"host" json:"host"`
	OtelInfo    OtelConfig        `mapstructure:"otel" json:"otel"`
	HomestaySrv HomestaySrvConfig `mapstructure:"homestay_srv" json:"homestay_srv"`
	UserSrv     UserSrvConfig     `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo     JWTConfig         `mapstructure:"jwt" json:"jwt"`
}
