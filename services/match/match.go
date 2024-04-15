package match

import (
	"context"

	db "github.com/corentings/chessbet/db/sqlc"
)

type IUseCase interface {
	GetMatchByID(ctx context.Context, id int32) (db.Match, error)
	GetMatches(ctx context.Context) ([]db.Match, error)
}
