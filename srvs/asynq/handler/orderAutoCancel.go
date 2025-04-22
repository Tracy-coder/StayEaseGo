package handler

import (
	"StayEaseGo/srvs/asynq/global"
	"StayEaseGo/srvs/asynq/model"
	order_model "StayEaseGo/srvs/order_srv/model"
	order_srv "StayEaseGo/srvs/order_srv/proto/gen"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
)

type OrderAutoCancelHandler struct {
}

func NewOrderAutoCancelHandler() *OrderAutoCancelHandler {
	return &OrderAutoCancelHandler{}
}

func (l *OrderAutoCancelHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p model.OrderAutoCancelMessage
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("OrderAutoCancelHandler: payload err:%v, payLoad:%+v", err, t.Payload())
	}
	log.Debugf("OrderAutoCancelHandler: payload:%+v", p)
	resp, err := global.OrderSrvClient.HomestayOrderDetail(ctx, &order_srv.HomestayOrderDetailReq{
		Sn: p.Sn,
	})
	if err != nil || resp.HomestayOrder == nil {
		return fmt.Errorf("OrderAutoCancelHandler: get order fail or order no exists err:%v, sn:%s ,HomestayOrder : %+v", err, p.Sn, resp.HomestayOrder)
	}

	if resp.HomestayOrder.TradeState == order_model.HomestayOrderTradeStateWaitPay {
		_, err := global.OrderSrvClient.UpdateHomestayOrderTradeState(ctx, &order_srv.UpdateHomestayOrderTradeStateReq{
			Sn:         p.Sn,
			TradeState: order_model.HomestayOrderTradeStateCancel,
		})
		if err != nil {
			return fmt.Errorf("OrderAutoCancelHandler: close order fail  err:%v, sn:%s ", err, p.Sn)
		}
	}

	return nil
}
