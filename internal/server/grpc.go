package server

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/viper"

	pbuser "gitlab.top.slotssprite.com/my/rpc-layout/api/helloworld/v1/user"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/conf"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/service"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/grpcx"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(user *service.UserService) *kgrpc.Server {
	var opts = []kgrpc.ServerOption{
		// NOTE: 添加中间件注意执行顺序！
		//grpcx.Validator()
		kgrpc.Middleware(grpcx.StartAt(), tracing.Server(), grpcx.Log(), grpcx.Recovery(), metadata.Server(), grpcx.RateLimit()),
		kgrpc.UnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
		kgrpc.StreamInterceptor(grpcprometheus.StreamServerInterceptor),
		//kgrpc.Options(grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)),
	}
	grpcHost := viper.GetString(conf.PathGrpcHost)
	grpcPort := viper.GetInt(conf.PathGrpcPort)
	if grpcHost != "" && grpcPort > 0 {
		opts = append(opts, kgrpc.Address(fmt.Sprintf("%s:%d", grpcHost, grpcPort)))
	}
	// NOTE: kgrpc.Timeout 暂不用设置，因为 unaryServerInterceptor 中并没有 select case <-ctx.Done()
	// NOTE: 设置了反倒会改变 context 的超时传递时间，一般情况 client 的 context 带有超时时间
	// NOTE: 正常情况下 client 调用超时是为了避免链路阻塞堆积，server 端继续处理请求也属正常
	// NOTE: 未自定义设置服务端timeout时，kratos框架默认设置为1秒，导致服务端调用时间过长或者链路较长时服务超时中断
	// NOTE: 所以此处更新timeout为0，即使用客户端调用传来的ctx中的超时控制
	opts = append(opts, kgrpc.Timeout(0))

	srv := kgrpc.NewServer(opts...)
	pbuser.RegisterUserServiceServer(srv, user)

	grpcprometheus.EnableHandlingTimeHistogram()
	grpcprometheus.Register(srv.Server)
	return srv
}
