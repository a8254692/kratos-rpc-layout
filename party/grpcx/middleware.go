package grpcx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/grpcx/group"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/grpcx/status"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/statusx"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/util"
	"runtime"
	"time"
)

// Validator ...
//func Validator() middleware.Middleware {
//	validate := validator.New()
//	zhT := zh.New()
//	uni := ut.New(zhT, zhT)
//	trans, _ := uni.GetTranslator("zh")
//	if err := zhTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
//		panic("validator middleware register translations error: " + err.Error())
//	}
//
//	return func(handler middleware.Handler) middleware.Handler {
//		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
//			if _err := validate.Struct(req); _err != nil {
//				if e, ok := (_err).(validator.ValidationErrors); ok {
//					var errBuffer bytes.Buffer
//					for k, v := range e.Translate(trans) {
//						errBuffer.WriteString(k)
//						errBuffer.WriteString(":")
//						errBuffer.WriteString(v)
//						errBuffer.WriteString(",")
//					}
//					errText := errBuffer.String()
//					if errBuffer.String() != "" {
//						errText = strings.TrimSuffix(errBuffer.String(), ",")
//					}
//					return nil, status.Error(ctx, errors.New(errText), statusx.StatusInvalidRequest)
//				}
//			}
//			return handler(ctx, req)
//		}
//	}
//}

// Recovery ...
func Recovery() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			defer func() {
				if _err := recover(); _err != nil {
					buf := make([]byte, 64<<10)
					n := runtime.Stack(buf, false)
					errStr := fmt.Sprintf("%v", _err)
					errInfo := map[string]interface{}{
						"error": errStr,
						"req":   fmt.Sprintf("%+v", req),
						"stack": fmt.Sprintf("%s", buf[:n]),
					}
					log.Context(ctx).Error(util.MustMarshalToString(errInfo))
					err = status.NewError(errors.New(errStr), statusx.StatusInternalServerError)
				}
			}()
			return handler(ctx, req)
		}
	}
}

// RateLimit ...
func RateLimit() middleware.Middleware {
	limiter := bbr.NewLimiter()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			done, e := limiter.Allow()
			if e != nil {
				return nil, status.NewError(e, statusx.StatusTooManyRequests)
			}
			reply, err = handler(ctx, req)
			done(ratelimit.DoneInfo{Err: err})
			return
		}
	}
}

// StartAt  ...
func StartAt() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			ctx = context.WithValue(ctx, "startAt", time.Now().UnixMilli())
			return handler(ctx, req)
		}
	}
}

// Log  ...
func Log() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				params = make(map[string]interface{})
				result = make(map[string]interface{})
			)
			if info, ok := transport.FromServerContext(ctx); ok {
				params["method"] = info.Operation()
				params["header"] = util.MustMarshalToString(info.RequestHeader())
				params["req"] = fmt.Sprint(req)
			}

			defer func() {
				st, _result := status.FromError(err)
				if st != nil {
					result["grpcCode"] = st.Code()
					result["error"] = st.Err().Error()
					if _result != nil {
						result["status"] = _result.Status
						result["msg"] = _result.Msg
						result["code"] = _result.Code
					}
				}

				var processTime int64
				if startAt, _ := ctx.Value("startAt").(int64); startAt > 0 {
					processTime = time.Now().UnixMilli() - startAt
				}
				log.Context(ctx).Log(log.LevelWarn,
					"@field", map[string]interface{}{
						"params":      params,
						"processTime": processTime,
						"result":      result,
					},
				)
			}()
			reply, err = handler(ctx, req)
			return
		}
	}
}

// ClientBreaker ...
func ClientBreaker() middleware.Middleware {
	gp := group.NewGroup(func() interface{} {
		// OPTIMIZE: NewBreaker Option ...
		return sre.NewBreaker()
	})
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			info, _ := transport.FromClientContext(ctx)
			breaker := gp.Get(info.Operation()).(circuitbreaker.CircuitBreaker)
			if err := breaker.Allow(); err != nil {
				breaker.MarkFailed()
				return nil, status.NewError(err, statusx.StatusTemporarilyUnavailable)
			}
			reply, err := handler(ctx, req)
			if err != nil && (kerrors.IsInternalServer(err) || kerrors.IsServiceUnavailable(err) || kerrors.IsGatewayTimeout(err)) {
				breaker.MarkFailed()
			} else {
				breaker.MarkSuccess()
			}
			return reply, err
		}
	}
}
