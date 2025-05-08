package handler

import (
	"StayEaseGo/pkg/xerr"
	"StayEaseGo/srvs/homestay_srv/config"
	"StayEaseGo/srvs/homestay_srv/global"
	"StayEaseGo/srvs/homestay_srv/model"
	pb "StayEaseGo/srvs/homestay_srv/proto/gen"
	"context"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type HomestaySever struct {
	svcCtx *ServiceContext
	pb.UnimplementedHomestayServer
}

type ServiceContext struct {
	Config    config.ServerConfig
	SqlClient *gorm.DB
}

func NewHomestaySever(svcCtx *ServiceContext) *HomestaySever {
	return &HomestaySever{svcCtx: svcCtx}
}

func (s *HomestaySever) HomestayDetail(ctx context.Context, req *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	var homestay model.Homestay
	result := s.svcCtx.SqlClient.Where("id = ?", req.ID).First(&homestay)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get homestay detail failed, id: %d", req.ID)
	}
	var homestayDetail pb.Homestay

	_ = copier.Copy(&homestayDetail, homestay)
	return &pb.HomestayDetailResp{
		Homestay: &homestayDetail,
	}, nil
}

func (s *HomestaySever) HomestayList(ctx context.Context, req *pb.HomestayListReq) (*pb.HomestayListResp, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{},
			},
		},
		"from": req.Offset,
		"size": req.Limit,
		"sort": []map[string]interface{}{
			{"_score": map[string]interface{}{"order": "desc"}},
		},
	}

	filters := query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"].([]map[string]interface{})

	if req.MinPrice != 0 {
		filters = append(filters, map[string]interface{}{
			"range": map[string]interface{}{
				"homestayPrice": map[string]interface{}{
					"gte": req.MinPrice,
				},
			},
		})
	}

	if req.MaxPrice != 0 {
		filters = append(filters, map[string]interface{}{
			"range": map[string]interface{}{
				"homestayPrice": map[string]interface{}{
					"lte": req.MaxPrice,
				},
			},
		})
	}

	if req.MinPeopleNum != 0 {
		filters = append(filters, map[string]interface{}{
			"range": map[string]interface{}{
				"peopleNum": map[string]interface{}{
					"gte": req.MinPeopleNum,
				},
			},
		})
	}

	if req.MaxPeopleNum != 0 {
		filters = append(filters, map[string]interface{}{
			"range": map[string]interface{}{
				"peopleNum": map[string]interface{}{
					"lte": req.MaxPeopleNum,
				},
			},
		})
	}

	if req.HomestayBusinessBossID != 0 {
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{
				"homestayBusinessBossID": req.HomestayBusinessBossID,
			},
		})
	}

	if req.RowType != 0 {
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{
				"rowType": req.RowType,
			},
		})
	}

	query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = filters

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ES_ERROR), "failed to marshal query: %s", err)
	}

	res, err := global.GlobalESClient.Search(
		global.GlobalESClient.Search.WithIndex("homestay"),
		global.GlobalESClient.Search.WithBody(strings.NewReader(string(queryJSON))),
	)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ES_ERROR), "failed to search: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "failed to decode error response")
		}
		log.Errorf("Search request failed: %s", e["error"])
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ES_ERROR), "search request failed with status %s", res.Status())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "failed to decode search response")
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	homestayList := make([]*pb.Homestay, 0, len(hits))
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		// log.Debug(source)
		homestay := &pb.Homestay{
			ID:                     int64(source["id"].(float64)),
			Title:                  source["title"].(string),
			SubTitle:               source["subTitle"].(string),
			Banner:                 source["banner"].(string),
			Info:                   source["info"].(string),
			PeopleNum:              int64(source["peopleNum"].(float64)),
			HomestayBusinessBossID: int64(source["homestayBusinessBossID"].(float64)),
			RowState:               int64(source["rowState"].(float64)),
			RowType:                int64(source["rowType"].(float64)),
			FoodInfo:               source["foodInfo"].(string),
			FoodPrice:              int64(source["foodPrice"].(float64)),
			HomestayPrice:          int64(source["homestayPrice"].(float64)),
			MarketHomestayPrice:    int64(source["marketHomestayPrice"].(float64)),
		}
		homestayList = append(homestayList, homestay)
	}

	return &pb.HomestayListResp{
		Homestays: homestayList,
	}, nil
}

func (s *HomestaySever) CreateHomestay(ctx context.Context, req *pb.CreateHomestayReq) (*pb.Empty, error) {
	var homestay model.Homestay
	_ = copier.Copy(&homestay, req.Homestay)
	result := s.svcCtx.SqlClient.Create(&homestay)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to create homestay: %s", result.Error)
	}
	return &pb.Empty{}, nil
}
