package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}

type ServerConfig struct {
	Name      string         `mapstructure:"name" json:"name"`
	Host      string         `mapstructure:"host" json:"host"`
	RedisInfo RedisConfig    `mapstructure:"redis" json:"redis"`
	OrderSrv  OrderSrvConfig `mapstructure:"order_srv" json:"order_srv"`
}
