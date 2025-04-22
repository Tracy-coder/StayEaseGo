package model

import (
	"time"

	"gorm.io/gorm"
)

type HomestayOrder struct {
	gorm.Model
	DelState               int64     `gorm:"del_state"`
	Version                int64     `gorm:"version"`               // 版本号
	Sn                     string    `gorm:"sn"`                    // 订单号
	UserId                 int64     `gorm:"user_id"`               // 下单用户id
	HomestayId             int64     `gorm:"homestay_id"`           // 民宿id
	Title                  string    `gorm:"title"`                 // 标题
	SubTitle               string    `gorm:"sub_title"`             // 副标题
	Cover                  string    `gorm:"cover"`                 // 封面
	Info                   string    `gorm:"info"`                  // 介绍
	PeopleNum              int64     `gorm:"people_num"`            // 容纳人的数量
	RowType                int64     `gorm:"row_type"`              // 售卖类型0：按房间出售 1:按人次出售
	NeedFood               bool      `gorm:"need_food"`             // false:不需要餐食 true:需要餐食
	FoodInfo               string    `gorm:"food_info"`             // 餐食标准
	FoodPrice              int64     `gorm:"food_price"`            // 餐食价格(分)
	HomestayPrice          int64     `gorm:"homestay_price"`        // 民宿价格(分)
	MarketHomestayPrice    int64     `gorm:"market_homestay_price"` // 民宿市场价格(分)
	HomestayBusinessBossID int64     `gorm:"homestay_user_id"`      // 店铺房东id
	LiveStartDate          time.Time `gorm:"live_start_date"`       // 开始入住日期
	LiveEndDate            time.Time `gorm:"live_end_date"`         // 结束入住日期
	LivePeopleNum          int64     `gorm:"live_people_num"`       // 实际入住人数
	TradeState             int64     `gorm:"trade_state"`           // -1: 已取消 0:待支付 1:未使用 2:已使用  3:已退款 4:已过期
	TradeCode              string    `gorm:"trade_code"`            // 确认码
	Remark                 string    `gorm:"remark"`                // 用户下单备注
	OrderTotalPrice        int64     `gorm:"order_total_price"`     // 订单总价格（餐食总价格+民宿总价格）(分)
	FoodTotalPrice         int64     `gorm:"food_total_price"`      // 餐食总价格(分)
	HomestayTotalPrice     int64     `gorm:"homestay_total_price"`  // 民宿总价格(分)
}

var HomestayOrderTradeStateCancel int64 = -1
var HomestayOrderTradeStateWaitPay int64 = 1
var HomestayOrderTradeStateWaitUse int64 = 2
var HomestayOrderTradeStateUsed int64 = 3
var HomestayOrderTradeStateRefund int64 = 4
var HomestayOrderTradeStateExpire int64 = 5

const (
	IsDeleted  = 1
	NotDeleted = 0
)
