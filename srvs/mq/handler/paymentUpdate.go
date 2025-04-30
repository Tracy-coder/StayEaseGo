package handler

import (
	"StayEaseGo/srvs/mq/global"
	"StayEaseGo/srvs/mq/model"
	order_model "StayEaseGo/srvs/order_srv/model"
	order_srv "StayEaseGo/srvs/order_srv/proto/gen"
	payment_model "StayEaseGo/srvs/payment_srv/model"
	"context"

	log "github.com/sirupsen/logrus"
)

func ThirdPaymentUpdatePayStatusNotifyHandler(payload model.ThirdPaymentUpdatePayStatusNotifyMessage) error {
	var newState int64
	log.Debug(payload)
	if payload.PayStatus == payment_model.ThirdPaymentPayTradeStateSuccess {
		newState = order_model.HomestayOrderTradeStateWaitUse
	} else if payload.PayStatus == payment_model.ThirdPaymentPayTradeStateRefund {
		newState = order_model.HomestayOrderTradeStateRefund
	} else {
		return nil
	}
	_, err := global.OrderSrvClient.UpdateHomestayOrderTradeState(context.Background(), &order_srv.UpdateHomestayOrderTradeStateReq{
		Sn:         payload.OrderSn,
		TradeState: newState,
	})
	return err
}
