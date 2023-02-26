package utils

import (
	"context"
	"github.com/dollarkillerx/zim/pkg/enums"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

func GRPCAuth(token string) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		md, ex := metadata.FromIncomingContext(ctx)
		if !ex {
			return ctx, errors.New("Authorization authentication failed")
		}

		PrintObject(md)

		rl := md.Get(enums.Token)
		if len(rl) == 0 {
			return ctx, errors.New("Authorization authentication failed")
		}

		if rl[0] != token {
			return ctx, errors.New("Authorization authentication failed")
		}

		return ctx, nil
	}
}

// AuthCredential 自定义认证
type AuthCredential struct {
	token string
}

func NewAuthCredential(token string) *AuthCredential {
	return &AuthCredential{token: token}
}

// GetRequestMetadata 实现自定义认证接口
func (c AuthCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		enums.Token: c.token,
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c AuthCredential) RequireTransportSecurity() bool {
	return false
}
