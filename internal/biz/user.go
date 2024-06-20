package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	pbuser "gitlab.top.slotssprite.com/my/rpc-layout/api/helloworld/v1/user"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/auth"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/grpcx/status"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/statusx"
)

// UserRepo db 层 crud
type UserRepo interface {
	GetAccountByPhone(ctx context.Context, phone string) (*ent.Account, error)
	GetPlayerById(ctx context.Context, id int64) (*ent.Player, error)
}

// UserUseCase ...
type UserUseCase struct {
	repo UserRepo
}

// NewUserUseCase ...
func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// GetAccountByPhone ...
func (s *UserUseCase) GetAccountByPhone(ctx context.Context, phone string) (*pbuser.Account, error) {
	info, err := s.repo.GetAccountByPhone(ctx, phone)
	if err != nil {
		log.Context(ctx).Error("GetAccountByPhone repo err:", err)
		return nil, err
	}
	pbAcc := new(pbuser.Account)

	log.Context(ctx).Info("GetAccountByPhone info", info)

	err = copier.Copy(&pbAcc, &info)
	log.Context(ctx).Info("GetAccountByPhone pbAcc", pbAcc)
	if err != nil {
		return nil, status.Error(ctx, err, statusx.StatusInternalServerError)
	}
	return pbAcc, nil
}

// GetPlayerById ...
func (s *UserUseCase) GetPlayerById(ctx context.Context) (*pbuser.Player, error) {
	uid := auth.GetCtxUser(ctx)
	if uid <= 0 {
		log.Context(ctx).Error(uid)
		return nil, errors.New("UserUseCase GetPlayerById auth is nil")
	}

	log.Context(ctx).Info(uid)

	info, err := s.repo.GetPlayerById(ctx, uid)
	if err != nil {
		return nil, err
	}

	pbUser := &pbuser.Player{}

	err = copier.Copy(&pbUser, &info)
	if err != nil {
		return nil, status.Error(ctx, err, statusx.StatusInternalServerError)
	}
	return pbUser, nil
}
