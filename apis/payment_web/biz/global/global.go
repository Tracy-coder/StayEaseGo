package global

import (
	"StayEaseGo/apis/payment_web/biz/config"
	order_srv "StayEaseGo/srvs/order_srv/proto/gen"
	payment_srv "StayEaseGo/srvs/payment_srv/proto/gen"
	user_srv "StayEaseGo/srvs/user_srv/proto/gen"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

const PaymentServiceConfigPath = "./apis/payment_web/biz/config/config_docker.yaml"

var (
	GlobalServerConfig config.ServerConfig
	GlobalConsulConfig config.ConsulConfig
)

var (
	PaymentSrvClient payment_srv.PaymentClient
	OrderSrvClient   order_srv.OrderClient
	UserSrvClient    user_srv.UserClient
)

var (
	WXPayClient *core.Client
)

const ThirdPaymentWxPay = "wxpay" // 第三方支付方式：微信支付
