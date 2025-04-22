package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers" json:"brokers"`
}
type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type ServerConfig struct {
	Name     string         `mapstructure:"name" json:"name"`
	Host     string         `mapstructure:"host" json:"host"`
	KafkaCfg KafkaConfig    `mapstructure:"kafka" json:"kafka"`
	UserSrv  UserSrvConfig  `mapstructure:"user_srv" json:"user_srv"`
	OrderSrv OrderSrvConfig `mapstructure:"order_srv" json:"order_srv"`
}
