package rpc

import (
	"context"
	usercore "github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore"
	"github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore/usercoreservice"
	"github.com/BlueGopher/tiktok-rpc/pkg/consts"
	"github.com/BlueGopher/tiktok-rpc/pkg/errno"
	"github.com/BlueGopher/tiktok-rpc/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var userCoreClient usercoreservice.Client

// initUser rpc usercore init client
func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.UserServiceName),
		provider.WithExportEndpoint(consts.ExprotEndpoint),
		provider.WithInsecure(),
	)
	c, err := usercoreservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),

		client.WithRPCTimeout(5*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
	)
	if err != nil {
		panic(err)
	}
	userCoreClient = c
}

// RegisterUser rpc : register user
func RegisterUser(ctx context.Context, request *usercore.DouyinUserRegisterRequest) error {
	resp, err := userCoreClient.RegisterUser(ctx, request)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// GetUserInfo rpc : get user info
func GetUserInfo(ctx context.Context, request *usercore.DouyinUserRequest) (*usercore.User, error) {
	resp, err := userCoreClient.GetUserInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.User, nil
}
