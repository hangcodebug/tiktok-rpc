package pack

import (
	"errors"
	usercore "github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore"
	"github.com/BlueGopher/tiktok-rpc/pkg/errno"
	"time"
)

func BuildBaseResponse(err error) *usercore.BaseResponse {
	if err != nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *usercore.BaseResponse {
	return &usercore.BaseResponse{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
