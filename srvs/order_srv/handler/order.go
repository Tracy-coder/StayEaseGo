package handler

import (
	"StayEaseGo/pkg/snowflake"
	"StayEaseGo/pkg/xerr"
	homestay "StayEaseGo/srvs/homestay_srv/proto/gen"
	mq "StayEaseGo/srvs/mq/model"
	"StayEaseGo/srvs/order_srv/config"
	"StayEaseGo/srvs/order_srv/global"
	"StayEaseGo/srvs/order_srv/model"
	pb "StayEaseGo/srvs/order_srv/proto/gen"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/pkg/errors"

	jobtype "StayEaseGo/srvs/asynq/model"

	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderSever struct {
	svcCtx *ServiceContext
	pb.UnimplementedOrderServer
}

type ServiceContext struct {
	Config            config.ServerConfig
	SqlClient         *gorm.DB
	HomestaySrvClient homestay.HomestayClient
}

func NewOrderSever(svcCtx *ServiceContext) *OrderSever {
	return &OrderSever{svcCtx: svcCtx}
}

func (s *OrderSever) CreateHomestayOrder(ctx context.Context, req *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	if req.LiveStartTime >= req.LiveEndTime {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "start time cannot be greater than or equal to homestay end time:start time: %s, end time: %s", time.Unix(req.LiveStartTime, 0), time.Unix(req.LiveEndTime, 0))
	}
	rpcResp, err := s.svcCtx.HomestaySrvClient.HomestayDetail(ctx, &homestay.HomestayDetailReq{ID: req.HomestayId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.RPCCALL_ERROR), "rpc HomestayDetail fail , homestayId : %d , err : %v", req.HomestayId, err)
	}
	if rpcResp == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "homestay not exists:ID: %d", req.HomestayId)
	}
	start := time.Unix(req.LiveStartTime, 0)
	end := time.Unix(req.LiveEndTime, 0)
	result := s.svcCtx.SqlClient.Model(&model.HomestayOrder{}).Where(
		"homestay_id = ? AND NOT (live_end_date < ? OR live_start_date > ?)",
		req.HomestayId, start, end,
	).First(&model.HomestayOrder{})
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "homestay is not available at this time:ID: %d", req.HomestayId)
	}
	order := new(model.HomestayOrder)
	order.Sn = snowflake.GenerateOrderID()
	order.HomestayId = req.HomestayId
	order.UserId = req.UserId
	order.Title = rpcResp.Homestay.Title
	order.SubTitle = rpcResp.Homestay.SubTitle
	if len(rpcResp.Homestay.Banner) > 0 {
		order.Cover = strings.Split(rpcResp.Homestay.Banner, ",")[0]
	}
	order.Info = rpcResp.Homestay.Info
	order.PeopleNum = rpcResp.Homestay.PeopleNum
	order.RowType = rpcResp.Homestay.RowType
	order.NeedFood = req.IsFood
	order.FoodInfo = rpcResp.Homestay.FoodInfo
	order.HomestayPrice = rpcResp.Homestay.HomestayPrice
	order.MarketHomestayPrice = rpcResp.Homestay.MarketHomestayPrice
	order.HomestayBusinessBossID = rpcResp.Homestay.HomestayBusinessBossID
	order.LiveStartDate = time.Unix(req.LiveStartTime, 0)
	order.LiveEndDate = time.Unix(req.LiveEndTime, 0)
	order.LivePeopleNum = req.LivePeopleNum
	order.TradeState = model.HomestayOrderTradeStateWaitPay
	order.Remark = req.Remark
	order.FoodPrice = rpcResp.Homestay.FoodPrice
	liveDays := int64(order.LiveEndDate.Sub(order.LiveStartDate).Seconds() / 86400)
	if order.NeedFood {
		order.FoodTotalPrice = liveDays * order.FoodPrice
	}
	order.HomestayTotalPrice = liveDays * order.HomestayPrice
	order.OrderTotalPrice = order.FoodTotalPrice + order.HomestayTotalPrice
	result = s.svcCtx.SqlClient.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	///// message queue to delay cancel
	payload, err := json.Marshal(jobtype.OrderAutoCancelMessage{
		Sn: order.Sn,
	})
	if err != nil {
		return nil, err
	}
	global.AsynqClient.Enqueue(asynq.NewTask(jobtype.JobTypeOrderAutoCancelNotify, payload), asynq.MaxRetry(3), asynq.ProcessAt(time.Now().Add(time.Hour*30)))
	log.Debugf("Enqueue order auto cancel task:sn: %s", order.Sn)
	return &pb.CreateHomestayOrderResp{Sn: order.Sn}, nil
}

func (s *OrderSever) HomestayOrderDetail(ctx context.Context, req *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {
	var order model.HomestayOrder
	result := s.svcCtx.SqlClient.Where(&model.HomestayOrder{Sn: req.Sn, DelState: model.NotDeleted}).First(&order)
	if result.RowsAffected == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "order not exists, sn: %s", req.Sn)
	}
	var respOrder pb.HomestayOrder
	_ = copier.Copy(&respOrder, order)
	return &pb.HomestayOrderDetailResp{HomestayOrder: &respOrder}, nil
}

func (s *OrderSever) UserHomestayOrderList(ctx context.Context, req *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	var order model.HomestayOrder
	order.UserId = req.UserId
	if req.TraderState != 0 {
		order.TradeState = req.TraderState
	}
	var res []model.HomestayOrder
	s.svcCtx.SqlClient.Where(&order).Offset(int(req.LastId)).Limit(int(req.PageSize)).Find(&res)
	respOrder := make([]*pb.HomestayOrder, len(res))
	for i, v := range res {
		respOrder[i] = new(pb.HomestayOrder)
		_ = copier.Copy(respOrder[i], v)
	}
	return &pb.UserHomestayOrderListResp{List: respOrder}, nil
}

func (s *OrderSever) UpdateHomestayOrderTradeState(ctx context.Context, req *pb.UpdateHomestayOrderTradeStateReq) (*pb.UpdateHomestayOrderTradeStateResp, error) {
	var order model.HomestayOrder
	log.Debugf("update order trade state:sn: %s, trade state: %d", req.Sn, req.TradeState)
	tx := s.svcCtx.SqlClient.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 悲观锁
	result := tx.Where(&model.HomestayOrder{Sn: req.Sn, DelState: model.NotDeleted}).First(&order).Clauses(clause.Locking{Strength: "UPDATE"}).Select("*")
	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "order not exists, sn: %s", req.Sn)
	}

	if !checkTradeState(order.TradeState, req.TradeState) {
		tx.Rollback()
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "update order trade state error")
	}

	res := tx.Model(&model.HomestayOrder{}).Where("sn = ?", req.Sn).Update("trade_state", req.TradeState)
	if res.Error != nil {
		tx.Rollback()
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), res.Error.Error())
	}
	tx.Commit()

	if req.TradeState == model.HomestayOrderTradeStateWaitUse {
		// todo:message queue
		data, err := json.Marshal(mq.OrderSuccessNotifyUserMessage{
			Sn:              order.Sn,
			UserId:          order.UserId,
			Title:           order.Title,
			OrderTotalPrice: order.OrderTotalPrice,
			LiveStartDate:   order.LiveStartDate.Unix(),
			LiveEndDate:     order.LiveEndDate.Unix(),
		})
		if err != nil {
			return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "json marshal failed")
		}
		err = global.KafkaProducer.WriteMessages(ctx,
			kafka.Message{
				Value: data,
			})
		if err != nil {
			log.Error("send message failed: ", err)
			return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "send kafka message failed")
		}
	}
	return &pb.UpdateHomestayOrderTradeStateResp{
		ID:              int64(order.ID),
		UserId:          order.UserId,
		Sn:              order.Sn,
		LiveStartDate:   order.LiveStartDate.Unix(),
		LiveEndDate:     order.LiveEndDate.Unix(),
		TradeCode:       order.TradeCode,
		OrderTotalPrice: order.OrderTotalPrice,
		Title:           order.Title,
	}, nil
}

func checkTradeState(oldTradeState, newTradeState int64) bool {

	if newTradeState == model.HomestayOrderTradeStateWaitPay {
		return false
	}

	if newTradeState == model.HomestayOrderTradeStateCancel {
		if oldTradeState == model.HomestayOrderTradeStateWaitPay {
			return true
		}

	} else if newTradeState == model.HomestayOrderTradeStateWaitUse {
		if oldTradeState == model.HomestayOrderTradeStateWaitPay {
			return true
		}
	} else if newTradeState == model.HomestayOrderTradeStateUsed {
		if oldTradeState == model.HomestayOrderTradeStateWaitUse {
			return true
		}
	} else if newTradeState == model.HomestayOrderTradeStateRefund {
		if oldTradeState == model.HomestayOrderTradeStateWaitUse {
			return true
		}
	} else if newTradeState == model.HomestayOrderTradeStateExpire {
		if oldTradeState == model.HomestayOrderTradeStateWaitUse {
			return true
		}
	}

	return false
}
