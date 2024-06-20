package server

import (
	"fmt"
	tHttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/viper"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/env"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/httpx"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/monitor"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/util"
)

// NewMonitorHTTPServer ...
func NewMonitorHTTPServer() (*httpx.MonitorServer, error) {
	opts := make([]tHttp.ServerOption, 0)
	httpHost := viper.GetString(env.PathMonitorHttpHost)
	httpPort := viper.GetInt(env.PathMonitorHttpPort)
	if httpPort < 0 {
		var err error
		if httpPort, err = util.GetFreePort(); err != nil {
			return nil, err
		}
	}
	if httpHost != "" && httpPort > 0 {
		opts = append(opts, tHttp.Address(fmt.Sprintf("%s:%d", httpHost, httpPort)))
	}
	httpSrv := tHttp.NewServer(opts...)
	httpSrv.HandlePrefix("/", monitor.DefaultServeMux)
	return (*httpx.MonitorServer)(httpSrv), nil
}
