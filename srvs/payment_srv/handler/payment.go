package handler

import (
	mq "StayEaseGo/srvs/mq/model"
	"StayEaseGo/srvs/payment_srv/config"
	"StayEaseGo/srvs/payment_srv/global"
	"StayEaseGo/srvs/payment_srv/model"
	pb "StayEaseGo/srvs/payment_srv/proto/gen"
	"StayEaseGo/srvs/pkg/snowflake"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type PaymentSever struct {
	svcCtx *ServiceContext
	pb.UnimplementedPaymentServer
}

type ServiceContext struct {
	Config    config.ServerConfig
	SqlClient *gorm.DB
}

func NewPaymentSever(svcCtx *ServiceContext) *PaymentSever {
	return &PaymentSever{svcCtx: svcCtx}
}

func (s *PaymentSever) CreatePayment(ctx context.Context, req *pb.CreatePaymentReq) (*pb.CreatePaymentResp, error) {
	payment := new(model.Payment)
	payment.Sn = snowflake.GeneratePaymentID()
	payment.UserId = req.UserId
	payment.PayMode = req.PayModel
	payment.PayTotal = req.PayTotal
	payment.PayStatus = model.ThirdPaymentPayTradeStateWait
	payment.ServiceType = req.ServiceType
	payment.OrderSn = req.OrderSn
	res := s.svcCtx.SqlClient.Create(payment)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("create payment failed")
	}
	return &pb.CreatePaymentResp{Sn: payment.Sn}, nil
}

func (s *PaymentSever) GetPaymentBySn(ctx context.Context, req *pb.GetPaymentBySnReq) (*pb.GetPaymentBySnResp, error) {
	var payment model.Payment
	result := s.svcCtx.SqlClient.Where(&model.Payment{Sn: req.Sn, DelState: model.NotDeleted}).First(&payment)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("payment not exists")
	}
	var respPayment pb.PaymentDetail
	_ = copier.Copy(&respPayment, payment)
	respPayment.PayTime = payment.PayTime.Unix()
	respPayment.CreateAt = payment.CreatedAt.Unix()
	respPayment.UpdateAt = payment.UpdatedAt.Unix()
	return &pb.GetPaymentBySnResp{PaymentDetail: &respPayment}, nil
}

func (s *PaymentSever) UpdateTradeState(ctx context.Context, req *pb.UpdateTradeStateReq) (*pb.UpdateTradeStateResp, error) {
	var payment model.Payment
	result := s.svcCtx.SqlClient.Where(&model.Payment{Sn: req.Sn, DelState: model.NotDeleted}).First(&payment)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("payment not exists")
	}
	if req.PayStatus == model.ThirdPaymentPayTradeStateFAIL || req.PayStatus == model.ThirdPaymentPayTradeStateSuccess {
		if payment.PayStatus != model.ThirdPaymentPayTradeStateWait {
			return nil, fmt.Errorf("payment status error")
		}
	}
	if req.PayStatus == model.ThirdPaymentPayTradeStateRefund {
		if payment.PayStatus != model.ThirdPaymentPayTradeStateSuccess {
			return nil, fmt.Errorf("payment status error")
		}
	}
	if req.PayStatus == model.ThirdPaymentPayTradeStateWait {
		return nil, fmt.Errorf("payment status error")
	}
	payment.PayStatus = req.PayStatus
	payment.TradeState = req.TradeState
	payment.TransactionId = req.TransactionId
	payment.TradeType = req.TradeType
	payment.TradeStateDesc = req.TradeStateDesc
	payment.PayTime = time.Unix(req.PayTime, 0)
	res := s.svcCtx.SqlClient.Save(&payment)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("update payment failed")
	}
	//todo: message queue : update order state
	data, err := json.Marshal(mq.ThirdPaymentUpdatePayStatusNotifyMessage{
		OrderSn:   payment.OrderSn,
		PayStatus: req.PayStatus,
	})
	if err != nil {
		return nil, fmt.Errorf("json marshal failed")
	}
	err = global.KafkaProducer.WriteMessages(ctx,
		kafka.Message{
			Value: data,
		})
	if err != nil {
		return nil, fmt.Errorf("send message failed")
	}
	return &pb.UpdateTradeStateResp{}, nil
}
