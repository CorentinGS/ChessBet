package tournament

import (
	"context"
	"log/slog"
	"time"

	db "github.com/corentings/chessbet/db/sqlc"
	"github.com/corentings/chessbet/pkg/lichess"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UseCase struct {
	q *db.Queries
}

func NewUseCase(dbConn *pgxpool.Pool) IUseCase {
	q := db.New(dbConn)

	return &UseCase{q: q}
}

func (u *UseCase) GetTournamentByID(_ context.Context, _ int32) (db.Tournament, error) {
	return db.Tournament{}, nil
}

func (u *UseCase) GetTournaments(_ context.Context) ([]db.Tournament, error) {
	return []db.Tournament{}, nil
}

func (u *UseCase) GetTournamentsInProgress(ctx context.Context) ([]db.Tournament, error) {
	return u.q.GetTournamentInProgress(ctx)
}

func (u *UseCase) CreateTournament(ctx context.Context, tournament db.CreateTournamentParams) (db.Tournament, error) {
	return u.q.CreateTournament(ctx, tournament)
}

func (u *UseCase) CreateTournamentFromLichessID(ctx context.Context, lichessID string) (db.Tournament, error) {
	const hoursInDay = 24

	broadcast, err := lichess.GetBroadcast(ctx, lichessID)
	if err != nil {
		return db.Tournament{}, err
	}

	tournament, err := u.q.CreateTournament(ctx, db.CreateTournamentParams{
		Name:                broadcast.Tour.Name,
		LichessTournamentID: lichessID,
		StartDate:           time.Now(),
		EndDate:             time.Now().Add(time.Hour * hoursInDay),
	})
	if err != nil {
		slog.Error("Failed to create tournament from lichess ID", slog.String("lichessID", lichessID), slog.String("error", err.Error()))
		return db.Tournament{}, err
	}

	return tournament, nil
}
