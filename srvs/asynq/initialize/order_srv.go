package initialize

import (
	"net"
	"strconv"

	"StayEaseGo/srvs/asynq/global"
	pb "StayEaseGo/srvs/order_srv/proto/gen"

	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitOrderSrv() {
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
	services, _, err := consulClient.Health().Service(global.GlobalServerConfig.OrderSrv.Name, "StayEaseGo", true, nil)
	if err != nil {
		log.Fatalf("Get user service failed: %s", err)
	}
	log.Debug(services)
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	grpcConn, _ := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcClient := pb.NewOrderClient(grpcConn)
	global.OrderSrvClient = grpcClient
}
