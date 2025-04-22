package model

type ESHomestay struct {
	ID                     int64   `json:"id"`
	Title                  string  `json:"title"`
	SubTitle               string  `json:"subTitle"`
	Banner                 string  `json:"banner"`
	Info                   string  `json:"info"`
	PeopleNum              int64   `json:"peopleNum"`
	HomestayBusinessBossID int64   `json:"homestayBusinessBossID"`
	RowState               int64   `json:"rowState"`            //0:下架 1:上架
	RowType                int64   `json:"rowType"`             //售卖类型0：按房间出售 1:按人次出售
	FoodInfo               string  `json:"foodInfo"`            //餐食标准
	FoodPrice              float64 `json:"foodPrice"`           //餐食价格
	HomestayPrice          float64 `json:"homestayPrice"`       //民宿价格
	MarketHomestayPrice    float64 `json:"marketHomestayPrice"` //民宿市场价格
}

func (ESHomestay) GetIndexName() string {
	return "homestay"
}
