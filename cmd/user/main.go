package main

import (
	"github.com/BlueGopher/tiktok-rpc/cmd/user/dal"
	usercore "github.com/BlueGopher/tiktok-rpc/kitex_gen/userCore/usercoreservice"
	"github.com/BlueGopher/tiktok-rpc/pkg/consts"
	"github.com/BlueGopher/tiktok-rpc/pkg/mw"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-etcd"
	"gorm.io/plugin/opentelemetry/provider"
	"net"
)

func Init() {
	dal.Init()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	TCPAddr, err := net.ResolveTCPAddr(consts.TCP, consts.UserServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	// 链路追踪
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.UserServiceName),
		provider.WithExportEndpoint(consts.ExprotEndpoint),
		provider.WithInsecure(),
	)
	svr := usercore.NewServer(new(UserCoreServiceImpl),
		server.WithServiceAddr(TCPAddr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: consts.UserServiceName,
		}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
