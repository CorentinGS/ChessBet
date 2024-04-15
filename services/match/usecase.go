package match

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

func (u *UseCase) GetMatchByID(_ context.Context, _ int32) (db.Match, error) {
	return db.Match{}, nil
}

func (u *UseCase) GetMatches(_ context.Context) ([]db.Match, error) {
	return []db.Match{}, nil
}
