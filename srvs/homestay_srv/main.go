package main

import (
	"StayEaseGo/srvs/homestay_srv/initialize"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	proto "StayEaseGo/srvs/homestay_srv/proto/gen"

	"StayEaseGo/pkg/addr"
	"StayEaseGo/pkg/interceptor"
	"StayEaseGo/srvs/homestay_srv/global"
	"StayEaseGo/srvs/homestay_srv/handler"

	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "address")
	Port := flag.Int("port", 0, "post")

	// initialization
	initialize.InitLogger()
	initialize.InitConfig()
	log.Debug(global.GlobalServerConfig)
	initialize.InitSql()
	initialize.InitES()
	flag.Parse()
	log.Info("ip: ", *IP)
	if *Port == 0 {
		*Port, _ = addr.GetFreePort()
	}

	log.Info("port: ", *Port)

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.LoggerInterceptor))
	ctx := &handler.ServiceContext{
		Config:    global.GlobalServerConfig,
		SqlClient: global.GlobalSqlClient,
	}
	proto.RegisterHomestayServer(server, handler.NewHomestaySever(ctx))
	reflection.Register(server)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	// Registration Service Health Check
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// service registry
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.GlobalConsulConfig.Host,
		global.GlobalConsulConfig.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// Generate corresponding inspection objects
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", global.GlobalServerConfig.Host, *Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	// Generate registration object
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.GlobalServerConfig.Name
	serviceID := uuid.NewV4().String()

	registration.ID = serviceID
	registration.Port = *Port
	registration.Tags = []string{"StayEaseGo", "homestay", "srv"}
	registration.Address = ""
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	// receive termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("sign out failed")
	} else {
		zap.S().Info("sign out success")
	}
}
