package data

import (
	"context"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/account"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/player"

	"gitlab.top.slotssprite.com/my/rpc-layout/internal/biz"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent"
)

type UserImp struct {
	data *Data
}

// NewUserImp .
func NewUserImp(data *Data) biz.UserRepo {
	return &UserImp{
		data: data,
	}
}

func (s *UserImp) GetAccountByPhone(ctx context.Context, phone string) (*ent.Account, error) {
	ps, err := s.data.db.Account.Query().Where(account.Phone(phone)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *UserImp) GetPlayerById(ctx context.Context, id int64) (*ent.Player, error) {
	ps, err := s.data.db.Player.Query().Where(player.PlayerID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ps, nil
}
