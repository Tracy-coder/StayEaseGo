// Code generated by hertz generator.

package main

import (
	"StayEaseGo/apis/homestay_web/biz/global"
	"StayEaseGo/apis/homestay_web/biz/initialize"
	"StayEaseGo/pkg/addr"
	"flag"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "address")
	Port := flag.Int("port", 0, "post")
	flag.Parse()
	if *Port == 0 {
		*Port, _ = addr.GetFreePort()
	}
	initialize.InitLogger()
	initialize.InitConfig()
	log.Debug(global.GlobalConsulConfig)
	initialize.InitUserSrv()
	initialize.InitHomestaySrv()
	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", *IP, *Port)))

	register(h)
	h.Spin()
}
