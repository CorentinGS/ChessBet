package tournament

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

func (u *UseCase) GetTournamentByID(ctx context.Context, id int32) (db.Tournament, error) {
	return db.Tournament{}, nil
}

func (u *UseCase) GetTournaments(ctx context.Context) ([]db.Tournament, error) {
	return []db.Tournament{}, nil
}
