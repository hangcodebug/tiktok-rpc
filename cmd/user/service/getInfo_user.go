package service

import (
	"context"
	"github.com/BlueGopher/tiktok-rpc/cmd/user/dal/db"
	"github.com/BlueGopher/tiktok-rpc/cmd/user/pack"
	usercore "github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore"
)

type GetUserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

func (s GetUserInfoService) GetUserInfo(req *usercore.DouyinUserRequest) (*usercore.User, error) {
	modelUser, err := db.GetUserInfo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.User(modelUser), nil
}
