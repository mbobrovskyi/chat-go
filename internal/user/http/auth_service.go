package http

import (
	"chat/internal/common/domain"
	userdomain "chat/internal/user/domain"
	"context"
)

type AuthService interface {
	SignIn(ctx context.Context, username, password string) (*userdomain.Token, error)
	SignUp(ctx context.Context, newUser domain.User, password string) (*userdomain.Token, error)
	SignOut(ctx context.Context, userID uint64, token string) error
	GetSession(ctx context.Context, token string) (*domain.Session, error)
}
