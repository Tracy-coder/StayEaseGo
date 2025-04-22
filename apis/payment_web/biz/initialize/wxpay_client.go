package initialize

import (
	"StayEaseGo/apis/payment_web/biz/global"
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func InitWXPayClient() {

	mchPrivateKey, err := utils.LoadPrivateKey(global.GlobalServerConfig.WxPayConf.PrivateKey)
	if err != nil {
		log.Fatalf("Load private key error: %s", err.Error())
	}

	ctx := context.Background()
	// Initialize the client with the merchant's private key, etc., and make it have the ability to automatically obtain WeChat payment platform certificates at regular intervals
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(global.GlobalServerConfig.WxPayConf.MchId, global.GlobalServerConfig.WxPayConf.SerialNo, mchPrivateKey, global.GlobalServerConfig.WxPayConf.APIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("Create client error: %s", err.Error())
	}
	global.WXPayClient = client

}
