package user

import (
	"context"

	db "github.com/corentings/chessbet/db/sqlc"
	"github.com/corentings/chessbet/pkg/oauth"
)

type IUseCase interface {
	GetUserByID(ctx context.Context, id int32) (db.User, error)
	GetUsers(ctx context.Context) ([]db.User, error)
	CreateUser(ctx context.Context, user db.CreateUserParams) (db.User, error)
	LoginOauthDiscord(ctx context.Context, oauth oauth.DiscordLogin) (string, error)
	RegisterOauth(ctx context.Context, oauth db.CreateUserParams) (db.User, error)
}
