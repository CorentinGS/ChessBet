package tournament

import (
	"context"

	db "github.com/corentings/chessbet/db/sqlc"
)

type IUseCase interface {
	GetTournamentByID(ctx context.Context, id int32) (db.Tournament, error)
	GetTournaments(ctx context.Context) ([]db.Tournament, error)
}
