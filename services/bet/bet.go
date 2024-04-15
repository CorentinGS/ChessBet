package bet

import (
	"context"

	db "github.com/corentings/chessbet/db/sqlc"
)

type IUseCase interface {
	GetBetByID(ctx context.Context, id int32) (db.Bet, error)
	GetBets(ctx context.Context) ([]db.Bet, error)
}
