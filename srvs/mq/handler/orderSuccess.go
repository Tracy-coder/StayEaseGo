package handler

import (
	"context"

	"StayEaseGo/srvs/mq/global"
	"StayEaseGo/srvs/mq/model"
	user_srv "StayEaseGo/srvs/user_srv/proto/gen"
)

func OrderSuccessNotifyUserHandler(payload model.OrderSuccessNotifyUserMessage) error {
	rpcResp, err := global.UserSrvClient.GetUserAuthByUserId(context.Background(), &user_srv.GetUserAuthByUserIDReq{UserID: payload.UserId})
	if err != nil || rpcResp == nil {
		return err
	}
	//todo 向用户发送消息
	// openID := rpcResp.UserAuth
	return nil
}
