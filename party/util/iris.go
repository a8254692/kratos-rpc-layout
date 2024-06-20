package util

import (
	"bytes"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

// GetInterface ...
func GetInterface(ctx iris.Context) (_interface string) {
	if len(ctx.Handlers()) > 0 {
		handle := strings.Split(ctx.HandlerName(), "/")
		_interface = StringBuilder(handle[2], ":", handle[len(handle)-2], ":", handle[len(handle)-1])
	}
	return
}

// GetProcessTime ...
func GetProcessTime(ctx iris.Context) (processTime int64) {
	startAt := ctx.Values().GetInt64Default("startAt", -1)
	if startAt > 0 {
		processTime = time.Now().UnixNano()/1000000 - startAt
	}
	return
}

// GetParamsString ...
func GetParamsString(ctx iris.Context) string {
	var paramsBuff bytes.Buffer
	ctx.Params().Visit(func(key string, value string) {
		paramsBuff.WriteString(key)
		paramsBuff.WriteString("=")
		paramsBuff.WriteString(value)
		paramsBuff.WriteString(",")
	})
	return paramsBuff.String()
}

// GetURLParamOfInclude ...
func GetURLParamOfInclude(ctx iris.Context, name string) (vs []string) {
	v := ctx.URLParam(name)
	if !strings.Contains(v, "include") {
		return
	}
	v = strings.ReplaceAll(v, "include", "")
	v = strings.ReplaceAll(v, "[", "")
	v = strings.ReplaceAll(v, "]", "")
	v = strings.TrimSpace(v)
	if v == "" {
		return
	}
	return strings.Split(v, ",")
}

// GetURLParamOfExclude ...
func GetURLParamOfExclude(ctx iris.Context, name string) (vs []string) {
	v := ctx.URLParam(name)
	if !strings.Contains(v, "exclude") {
		return
	}
	v = strings.ReplaceAll(v, "exclude", "")
	v = strings.ReplaceAll(v, "[", "")
	v = strings.ReplaceAll(v, "]", "")
	v = strings.TrimSpace(v)
	if v == "" {
		return
	}
	return strings.Split(v, ",")
}

// ContainInAndExclude ...
func ContainInAndExclude(v string) bool {
	return strings.Contains(v, "include") || strings.Contains(v, "exclude")
}

// WithDefaultRemoteAddrHeaders ...
func WithDefaultRemoteAddrHeaders(headerNames ...string) iris.Configurator {
	return func(app *iris.Application) {
		if len(headerNames) == 0 {
			app.Configure(iris.WithRemoteAddrHeader("X-Forwarded-For"), iris.WithRemoteAddrHeader("X-Real-Ip"))
			return
		}
		for _, headerName := range headerNames {
			app.Configure(iris.WithRemoteAddrHeader(headerName))
		}
	}
}
