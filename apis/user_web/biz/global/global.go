package global

import (
	"StayEaseGo/apis/user_web/biz/config"
	pb "StayEaseGo/srvs/user_srv/proto/gen"
)

const UserServiceConfigPath = "./apis/user_web/biz/config/config_docker.yaml"

var UserAuthTypeSystem string = "system"
var UserAuthTypeSmallWX string = "wxMini"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	UserSrvClient pb.UserClient
)
