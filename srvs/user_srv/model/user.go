package model

import "time"

type User struct {
	ID         int64     `gorm:"column:id;primaryKey"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
	DeleteTime time.Time `gorm:"column:delete_time;default:null"`
	DelState   int8      `gorm:"column:del_state"`
	Version    int8      `gorm:"column:version;default:1;comment:版本号"`
	Mobile     string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:手机号"`
	Password   string    `gorm:"column:password;not null;comment:密码"`
	Nickname   string    `gorm:"column:nickname;not null;comment:昵称"`
	Sex        int8      `gorm:"column:sex;default:null;comment:性别 0:男 1:女"`
	Avatar     string    `gorm:"column:avatar;comment:头像"`
	Info       string    `gorm:"column:info;comment:个人简介"`
}

const (
	IsDeleted  = 1
	NotDeleted = 0
)

type UserAuth struct {
	ID         int64     `gorm:"column:id;primaryKey"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
	DeleteTime time.Time `gorm:"column:delete_time;default:null"`
	DelState   int64     `gorm:"column:del_state"`
	Version    int64     `gorm:"column:version;comment:版本号"`
	UserId     int64     `gorm:"column:user_id"`
	AuthKey    string    `gorm:"column:auth_key;comment:平台唯一id"`
	AuthType   string    `gorm:"column:auth_type;comment:平台类型"`
}
