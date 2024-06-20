package status

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pbresult "gitlab.top.slotssprite.com/my/rpc-layout/api/helloworld/v1/result"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/runtimex"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/statusx"
	"google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

const (
	// BusinessError 自定义的业务错误 code
	BusinessError codes.Code = 101

	// _code HACK
	_code = "11000001"
)

// NewError ...
func NewError(err error, status statusx.Status) error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	st, _ := gstatus.New(BusinessError, errMsg).WithDetails(&pbresult.Result{
		Status: status.String(),
		Msg:    statusx.GetMsg(status),
		Code:   _code,
	})
	return st.Err()
}

// Error ...
func Error(ctx context.Context, err error, status statusx.Status) error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	st, _ := gstatus.New(BusinessError, errMsg).WithDetails(&pbresult.Result{
		Status: status.String(),
		Msg:    statusx.GetMsg(status),
		Code:   _code,
	})
	logger := log.NewHelper(log.With(log.GetLogger(), "@errorf", runtimex.Caller(5)))
	logger.WithContext(ctx).Error(st.Err())
	return st.Err()
}

// ErrorWithMsg ...
func ErrorWithMsg(ctx context.Context, err error, status statusx.Status, msg string) error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	st, _ := gstatus.New(BusinessError, errMsg).WithDetails(&pbresult.Result{
		Status: status.String(),
		Msg:    msg,
		Code:   _code,
	})
	logger := log.NewHelper(log.With(log.GetLogger(), "@errorf", runtimex.Caller(5)))
	logger.WithContext(ctx).Error(st.Err())
	return st.Err()
}

// FromError ...
func FromError(err error) (*gstatus.Status, *pbresult.Result) {
	if err == nil {
		return nil, nil
	}
	if st, ok := gstatus.FromError(err); ok {
		for _, detail := range st.Details() {
			if _result, ok := detail.(*pbresult.Result); ok {
				return st, _result
			}
		}
		return st, nil
	}
	return nil, nil
}
