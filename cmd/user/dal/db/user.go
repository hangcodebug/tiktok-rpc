package db

import (
	"context"
	"github.com/BlueGopher/tiktok-rpc/pkg/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// RegisterUser 创建用户
func RegisterUser(ctx context.Context, user User) error {
	return DB.WithContext(ctx).Create(&user).Error
}

// GetUserInfo 得到用户信息
func GetUserInfo(ctx context.Context, userID int64) (*User, error) {
	//user := make([]*User, 0)
	user := &User{}
	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
