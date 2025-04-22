package initialize

import (
	"net"
	"strconv"

	"StayEaseGo/srvs/mq/global"

	"StayEaseGo/srvs/pkg/addr"

	"github.com/bytedance/sonic"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitConfig to init consul config server
func InitConfig() {
	v := viper.New()
	v.SetConfigFile(global.MQServiceConfigPath)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&global.GlobalConsulConfig); err != nil {
		log.Fatalf("unmarshal err failed: %s", err.Error())
	}
	log.Infof("Config Info: %v", global.GlobalConsulConfig)

	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		global.GlobalConsulConfig.Host,
		strconv.Itoa(global.GlobalConsulConfig.Port))
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		log.Fatalf("new consul client failed: %s", err.Error())
	}
	content, _, err := consulClient.KV().Get(global.GlobalConsulConfig.Key, nil)
	if err != nil {
		log.Fatalf("consul kv failed: %s", err.Error())
	}
	// log.Debug(&global.GlobalServerConfig, content)
	err = sonic.Unmarshal(content.Value, &global.GlobalServerConfig)
	if err != nil {
		log.Fatalf("sonic unmarshal config failed: %s", err.Error())
	}

	if global.GlobalServerConfig.Host == "" {
		global.GlobalServerConfig.Host, err = addr.GetLocalIPv4Address()
		if err != nil {
			log.Fatalf("get localIpv4Addr failed:%s", err.Error())
		}
	}
}
