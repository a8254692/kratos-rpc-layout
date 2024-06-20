package auth

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/metadata"

	"gitlab.top.slotssprite.com/my/rpc-layout/party/util"
)

// GetCtxUser ...
func GetCtxUser(ctx context.Context) int64 {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return 0
	}

	fmt.Println("==============1=============", md)

	userIdStr := md.Get("x-md-global-user-id")
	if userIdStr == "" {
		return 0
	}

	return int64(util.StringToInt(userIdStr))
}
