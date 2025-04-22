package model

type ThirdPaymentUpdatePayStatusNotifyMessage struct {
	OrderSn   string
	PayStatus int64
}

type OrderSuccessNotifyUserMessage struct {
	Sn              string
	UserId          int64
	Title           string
	OrderTotalPrice int64
	LiveStartDate   int64
	LiveEndDate     int64
}
