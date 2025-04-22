package model

import (
	"StayEaseGo/srvs/homestay_srv/global"
	"bytes"
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Homestay struct {
	ID                     int64   `gorm:"column:id;primary_key"`
	Title                  string  `gorm:"column:title;not null"`
	SubTitle               string  `gorm:"subTitle"`
	Banner                 string  `gorm:"banner"`
	Info                   string  `gorm:"info"`
	PeopleNum              int64   `gorm:"peopleNum"`
	HomestayBusinessBossID int64   `gorm:"homestayBusinessBossID"`
	RowState               int64   `gorm:"column:rowState;not null;default:1"` //0:下架 1:上架
	RowType                int64   `gorm:"rowType"`                            //售卖类型0：按房间出售 1:按人次出售
	FoodInfo               string  `gorm:"foodInfo"`                           //餐食标准
	FoodPrice              float64 `gorm:"foodPrice"`                          //餐食价格
	HomestayPrice          float64 `gorm:"homestayPrice"`                      //民宿价格
	MarketHomestayPrice    float64 `gorm:"marketHomestayPrice"`                //民宿市场价格
}

type HomestayComment struct {
	ID         int64   `gorm:"id"`
	HomestayID int64   `gorm:"homestayId"`
	Content    string  `gorm:"content"`
	Star       float64 `gorm:"star"`
	UserID     int64   `gorm:"userId"`
	Nickname   string  `gorm:"nickname"`
	Avatar     string  `gorm:"avatar"`
}

//	type HomestayBusiness struct {
//		ID        int64   `gorm:"id"`
//		Title     string  `gorm:"title"` //店铺名称
//		Info      string  `gorm:"info"`  //店铺介绍
//		Tags      string  `gorm:"tags"`  //标签，多个用“,”分割
//		Cover     string  `gorm:"cover"` //
//		Star      float64 `gorm:"star"`
//		IsFav     int64   `gorm:"isFav"`
//		HeaderImg string  `gorm:"headerImg"` //店招门头图片
//	}
type HomestayBusinessBoss struct {
	ID       int64  `gorm:"id"`
	UserID   int64  `gorm:"userId"`
	Nickname string `gorm:"nickname"`
	Avatar   string `gorm:"avatar"`
	Info     string `gorm:"info"` //房东介绍
	Rank     int64  `gorm:"rank"` //排名
}

func (h *Homestay) AfterCreate(tx *gorm.DB) (err error) {
	log.Info("homestay after create")
	homestay := &ESHomestay{
		ID:                     h.ID,
		Title:                  h.Title,
		SubTitle:               h.SubTitle,
		Banner:                 h.Banner,
		Info:                   h.Info,
		PeopleNum:              h.PeopleNum,
		HomestayBusinessBossID: h.HomestayBusinessBossID,
		RowState:               h.RowState,
		RowType:                h.RowType,
		FoodInfo:               h.FoodInfo,
		FoodPrice:              h.FoodPrice,
		HomestayPrice:          h.HomestayPrice,
		MarketHomestayPrice:    h.MarketHomestayPrice,
	}
	data, err := json.Marshal(homestay)
	if err != nil {
		return err
	}
	_, err = global.GlobalESClient.Index(ESHomestay{}.GetIndexName(), bytes.NewReader(data), global.GlobalESClient.Index.WithDocumentID(strconv.FormatInt(h.ID, 10)))
	return err
}

func (h *Homestay) AfterUpdate(tx *gorm.DB) (err error) {
	homestay := &ESHomestay{
		ID:                     h.ID,
		Title:                  h.Title,
		SubTitle:               h.SubTitle,
		Banner:                 h.Banner,
		Info:                   h.Info,
		PeopleNum:              h.PeopleNum,
		HomestayBusinessBossID: h.HomestayBusinessBossID,
		RowState:               h.RowState,
		RowType:                h.RowType,
		FoodInfo:               h.FoodInfo,
		FoodPrice:              h.FoodPrice,
		HomestayPrice:          h.HomestayPrice,
		MarketHomestayPrice:    h.MarketHomestayPrice,
	}
	data, err := json.Marshal(homestay)
	if err != nil {
		return err
	}
	_, err = global.GlobalESClient.Update(ESHomestay{}.GetIndexName(), strconv.Itoa(int(h.ID)), bytes.NewReader(data))
	return err
}

func (h *Homestay) AfterDelete(tx *gorm.DB) (err error) {
	_, err = global.GlobalESClient.Delete(ESHomestay{}.GetIndexName(), strconv.Itoa(int(h.ID)))
	return
}
