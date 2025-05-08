package initialize

import (
	"net"
	"strconv"

	homestay "StayEaseGo/srvs/homestay_srv/proto/gen"
	"StayEaseGo/srvs/order_srv/global"

	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitHomestaySrv() {
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		global.GlobalConsulConfig.Host,
		strconv.Itoa(global.GlobalConsulConfig.Port))
	log.Debug(global.GlobalConsulConfig)
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		log.Fatalf("Connect to consul failed: %s", err)
	}
	log.Debug(global.GlobalServerConfig)
	services, _, err := consulClient.Health().Service(global.GlobalServerConfig.HomestaySrv.Name, "StayEaseGo", true, nil)
	if err != nil {
		log.Fatalf("Get homestay service failed: %s", err)
	}
	log.Debug(services)
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	global.GlobalServerConfig.HomestaySrv.Addr = addr
	grpcConn, _ := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcClient := homestay.NewHomestayClient(grpcConn)
	global.HomestaySrvClient = grpcClient
}
