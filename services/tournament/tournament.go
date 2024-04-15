package tournament

import (
	"context"

	db "github.com/corentings/chessbet/db/sqlc"
)

type IUseCase interface {
	GetTournamentByID(ctx context.Context, id int32) (db.Tournament, error)
	GetTournaments(ctx context.Context) ([]db.Tournament, error)
	CreateTournament(ctx context.Context, tournament db.CreateTournamentParams) (db.Tournament, error)
	CreateTournamentFromLichessID(ctx context.Context, lichessID string) (db.Tournament, error)
	GetTournamentsInProgress(ctx context.Context) ([]db.Tournament, error)
}
