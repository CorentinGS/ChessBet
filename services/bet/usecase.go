package bet

import (
	"context"

	db "github.com/corentings/chessbet/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UseCase struct {
	q *db.Queries
}

func NewUseCase(dbConn *pgxpool.Pool) IUseCase {
	q := db.New(dbConn)

	return &UseCase{q: q}
}

func (u *UseCase) GetBetByID(_ context.Context, _ int32) (db.Bet, error) {
	return db.Bet{}, nil
}

func (u *UseCase) GetBets(_ context.Context) ([]db.Bet, error) {
	return nil, nil
}
