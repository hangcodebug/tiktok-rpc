package main

import (
	"context"
	"github.com/BlueGopher/tiktok-rpc/cmd/user/pack"
	"github.com/BlueGopher/tiktok-rpc/cmd/user/service"
	usercore "github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore"
	"github.com/BlueGopher/tiktok-rpc/pkg/errno"
)

// UserCoreServiceImpl implements the last service interface defined in the IDL.
type UserCoreServiceImpl struct{}

// GetUserInfo implements the UserCoreServiceImpl interface.
func (s *UserCoreServiceImpl) GetUserInfo(ctx context.Context, req *usercore.DouyinUserRequest) (resp *usercore.DouyinUserResponse, err error) {
	// TODO: Your code here...
	resp = new(usercore.DouyinUserResponse)

	// Param error
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewGetUserInfoService(ctx).GetUserInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	resp.User = user
	return resp, nil
}

// RegisterUser implements the UserCoreServiceImpl interface.
func (s *UserCoreServiceImpl) RegisterUser(ctx context.Context, req *usercore.DouyinUserRegisterRequest) (resp *usercore.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(usercore.DouyinUserRegisterResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	return resp, nil
}
