package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Sn             string    `gorm:"column:sn;not null"` // 流水单号
	DelState       int64     `gorm:"del_state"`
	Version        int64     `gorm:"version"`                      // 乐观锁版本号
	UserId         int64     `gorm:"user_id"`                      // 用户id
	PayMode        string    `gorm:"pay_mode"`                     // 支付方式 1:微信支付
	TradeType      string    `gorm:"trade_type"`                   // 第三方支付类型
	TradeState     string    `gorm:"trade_state"`                  // 第三方交易状态
	PayTotal       int64     `gorm:"pay_total"`                    // 支付总金额(分)
	TransactionId  string    `gorm:"transaction_id"`               // 第三方支付单号
	TradeStateDesc string    `gorm:"trade_state_desc"`             // 支付状态描述
	OrderSn        string    `gorm:"order_sn"`                     // 业务单号
	ServiceType    string    `gorm:"service_type"`                 // 业务类型
	PayStatus      int64     `gorm:"pay_status"`                   // 平台内交易状态   -1:支付失败 0:未支付 1:支付成功 2:已退款
	PayTime        time.Time `gorm:"column:pay_time;default:null"` // 支付成功时间
}

const (
	WXPAY = "wxpay"
)

const (
	NotDeleted = 0
)

var ThirdPaymentPayTradeStateFAIL int64 = -1   //支付失败
var ThirdPaymentPayTradeStateWait int64 = 0    //待支付
var ThirdPaymentPayTradeStateSuccess int64 = 1 //支付成功
var ThirdPaymentPayTradeStateRefund int64 = 2  //已退款
