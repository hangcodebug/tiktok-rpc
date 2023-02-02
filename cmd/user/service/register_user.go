package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/BlueGopher/tiktok-rpc/cmd/user/dal/db"
	usercore "github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore"
	"io"
)

type RegisterUserService struct {
	ctx context.Context
}

// NewRegisterUserService new RegisterUserService
func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{ctx: ctx}
}

func (s *RegisterUserService) RegisterUser(req *usercore.DouyinUserRegisterRequest) error {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.RegisterUser(s.ctx, db.User{
		Name:     req.Username,
		Password: password,
	})
}
