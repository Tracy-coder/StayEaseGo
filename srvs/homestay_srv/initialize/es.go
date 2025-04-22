package initialize

import (
	"StayEaseGo/srvs/homestay_srv/global"
	"StayEaseGo/srvs/homestay_srv/model"
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
)

func InitES() {
	var err error
	addr := fmt.Sprintf("http://%s:%d", global.GlobalServerConfig.EsInfo.Host, global.GlobalServerConfig.EsInfo.Port)
	global.GlobalESClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{addr},
		Username:  "elastic",
		Password:  "123456",
	})
	if err != nil {
		log.Fatalf("Init ES failed: %s", err.Error())
	}
	res, err := esapi.IndicesExistsRequest{
		Index: []string{model.ESHomestay{}.GetIndexName()},
	}.Do(context.Background(), global.GlobalESClient)
	if err != nil {
		log.Fatalf("Init ES failed: %s", err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == 404 {
		log.Info("ES index not exists, create now")
		_, err = esapi.IndicesCreateRequest{
			Index: model.ESHomestay{}.GetIndexName(),
		}.Do(context.Background(), global.GlobalESClient)
		if err != nil {
			log.Fatalf("Init ES failed: %s", err.Error())
		}
	} else if res.StatusCode != 200 {
		log.Fatalf("Init ES failed: check indices exists return code %d", res.StatusCode)
	}
}
