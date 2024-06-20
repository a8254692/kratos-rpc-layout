package service

import (
	"context"
	pbuser "gitlab.top.slotssprite.com/my/rpc-layout/api/helloworld/v1/user"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

// NewUserService ...
func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

// UserService ...
type UserService struct {
	uc *biz.UserUseCase
}

func (u UserService) GetAccount(ctx context.Context, req *pbuser.GetAccountReq) (*pbuser.Account, error) {
	return u.uc.GetAccountByPhone(ctx, req.Phone)
}

func (u UserService) GetPlayerById(ctx context.Context, empty *emptypb.Empty) (*pbuser.Player, error) {
	return u.uc.GetPlayerById(ctx)
}
