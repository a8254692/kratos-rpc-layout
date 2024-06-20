package bootstrap

import (
	"context"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/bootstrap/logx"
	"time"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/env"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/runtimex"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/util"
)

// NewLogger ...
func NewLogger() (logger klog.Logger, cleanup func()) {
	switch env.GetMode() {
	case env.ModeDevelop, env.ModeTest, env.ModeProduction:
		logger, cleanup = logx.NewLogrusLogger()
	case env.ModeLocal:
		logger, cleanup = logx.NewLogrusLogger()
	default:
		return nil, nil
	}

	localIP, _ := util.GetLocalIP()
	logger = klog.With(logger,
		"@system", env.GetServiceName(),
		"@version", env.GetServiceVersion(),
		"@source", localIP,
		"@caller", runtimex.Caller(4),
		"@traceId", tracing.TraceID(),
		"@spanId", tracing.SpanID(),
		"@timestamp", _timestamp(),
		"@date", klog.Timestamp("2006-01-02 15:04:05"),
	)
	return
}

// _timestamp ...
func _timestamp() klog.Valuer {
	return func(ctx context.Context) interface{} {
		return time.Now().Unix()
	}
}
