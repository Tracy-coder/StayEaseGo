package handler

import (
	"StayEaseGo/srvs/user_srv/model"
	pb "StayEaseGo/srvs/user_srv/proto/gen"
	"context"
	"fmt"

	"StayEaseGo/srvs/user_srv/config"

	"StayEaseGo/srvs/pkg/encrypt"

	"github.com/jinzhu/copier"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserSever struct {
	svcCtx *ServiceContext
	pb.UnimplementedUserServer
}

type ServiceContext struct {
	Config      config.ServerConfig
	RedisClient *redis.Client
	SqlClient   *gorm.DB
}

func NewUserSever(svcCtx *ServiceContext) *UserSever {
	return &UserSever{svcCtx: svcCtx}
}
func (s *UserSever) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	var user model.User
	result := s.svcCtx.SqlClient.Where(&model.User{Mobile: req.AuthKey, DelState: model.NotDeleted}).First(&user)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not exists,mobile:%s", req.AuthKey)
	}
	if ok := encrypt.BcryptCheck(req.Password, user.Password); !ok {
		return nil, fmt.Errorf("wrong password")
	}

	return &pb.LoginResp{UserID: user.ID}, nil
}

func (s *UserSever) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	res := s.svcCtx.SqlClient.Where(&model.User{Mobile: req.Mobile, DelState: model.NotDeleted}).First(&model.User{})
	if res.Error != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("mobile already exists:%s", req.Mobile)
	}
	user := new(model.User)
	s.svcCtx.SqlClient.Transaction(func(tx *gorm.DB) error {
		encryptedPassword, err := encrypt.BcryptEncrypt(req.Password)
		if err != nil {
			return err
		}
		user.Mobile = req.Mobile
		user.Password = encryptedPassword
		user.Nickname = req.Nickname

		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		userAuth := model.UserAuth{
			UserId:   user.ID,
			AuthKey:  req.AuthKey,
			AuthType: req.AuthType,
		}
		if err := tx.Create(&userAuth).Error; err != nil {
			return err
		}
		return nil
	})
	return &pb.RegisterResp{UserID: user.ID}, nil
}

func (s *UserSever) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	var user model.User
	result := s.svcCtx.SqlClient.Where(&model.User{ID: req.Id, DelState: model.NotDeleted}).First(&user)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not exists")
	}
	var respUser pb.User
	_ = copier.Copy(&respUser, user)
	return &pb.GetUserInfoResp{User: &respUser}, nil
}

func (s *UserSever) GetUserAuthByUserId(ctx context.Context, req *pb.GetUserAuthByUserIDReq) (*pb.GetUserAuthByUserIDResp, error) {
	var userAuth model.UserAuth
	result := s.svcCtx.SqlClient.Where(&model.UserAuth{UserId: req.UserID, AuthType: req.AuthType}).First(&userAuth)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not exists")
	}
	var respUserAuth pb.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)
	return &pb.GetUserAuthByUserIDResp{UserAuth: &respUserAuth}, nil
}

func (s *UserSever) GetUserAuthByAuthKey(ctx context.Context, req *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	var userAuth model.UserAuth
	result := s.svcCtx.SqlClient.Where(&model.UserAuth{AuthKey: req.AuthKey, AuthType: req.AuthType}).First(&userAuth)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not exists")
	}
	var respUserAuth pb.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)
	return &pb.GetUserAuthByAuthKeyResp{UserAuth: &respUserAuth}, nil
}
