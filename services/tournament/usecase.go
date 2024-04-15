package tournament

import (
	"context"
	"errors"
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
	broadcast, err := lichess.GetBroadcast(ctx, lichessID)
	if err != nil {
		return db.Tournament{}, err
	}

	tournament, err := u.createTournament(ctx, broadcast, lichessID)
	if err != nil {
		slog.Error("Failed to create tournament", slog.String("error", err.Error()))
		return db.Tournament{}, err
	}

	for _, round := range broadcast.Rounds {
		roundDetails, getRoundErr := lichess.GetRound(ctx, round.ID)
		if getRoundErr != nil {
			slog.Error("Failed to get round details", slog.String("error", getRoundErr.Error()))
			continue
		}

		u.createMatches(ctx, roundDetails.Games, roundDetails, tournament, round)
	}

	return tournament, nil
}

const millisecondsToSeconds = 1000

func (u *UseCase) createTournament(ctx context.Context, broadcast lichess.Broadcast, lichessID string) (db.Tournament, error) {
	if len(broadcast.Rounds) == 0 {
		return db.Tournament{}, errors.New("no rounds found in broadcast")
	}

	return u.q.CreateTournament(ctx, db.CreateTournamentParams{
		Name:                broadcast.Tour.Name,
		LichessTournamentID: lichessID,
		StartDate:           time.Unix(broadcast.Rounds[0].StartsAt/millisecondsToSeconds, 0),
		EndDate:             time.Unix(broadcast.Rounds[len(broadcast.Rounds)-1].StartsAt/millisecondsToSeconds, 0),
	})
}

func (u *UseCase) createMatches(ctx context.Context, games []lichess.Game, roundDetails lichess.RoundDetails, tournament db.Tournament, round lichess.Round) {
	const minPlayers = 2
	matchParams := make([]db.CreateMatchesParams, 0, len(games))

	for _, game := range games {
		if len(game.Players) < minPlayers {
			slog.Error("Game has less than 2 players", slog.String("gameID", game.ID))
			continue
		}

		whitePlayer, err := u.getOrCreatePlayer(ctx, db.CreatePlayerParams{
			Name:   game.Players[0].Name,
			Rating: int32(game.Players[0].Rating), // Convert int to int32
		})
		if err != nil {
			slog.Error("Failed to get or create white player", slog.String("playerName", game.Players[0].Name), slog.String("error", err.Error()))
			continue
		}

		blackPlayer, err := u.getOrCreatePlayer(ctx, db.CreatePlayerParams{
			Name:   game.Players[1].Name,
			Rating: int32(game.Players[1].Rating), // Convert int to int32
		})
		if err != nil {
			slog.Error("Failed to get or create black player", slog.String("playerName", game.Players[1].Name), slog.String("error", err.Error()))
			continue
		}

		matchDate := time.Unix(roundDetails.Round.StartsAt, 0) // Convert int64 to time.Time

		matchParams = append(matchParams, db.CreateMatchesParams{
			TournamentID:   &tournament.TournamentID,
			Player1ID:      whitePlayer.PlayerID,
			Player2ID:      blackPlayer.PlayerID,
			MatchDate:      matchDate, // Use the converted time.Time value
			RoundName:      roundDetails.Round.Name,
			LichessRoundID: round.ID,
		})
	}

	_, err := u.q.CreateMatches(ctx, matchParams)
	if err != nil {
		slog.Error("Failed to create matches", slog.String("error", err.Error()))
	}
}

func (u *UseCase) getOrCreatePlayer(ctx context.Context, player db.CreatePlayerParams) (db.Player, error) {
	existingPlayer, err := u.q.GetPlayerByName(ctx, player.Name)
	if err != nil {
		// Player doesn't exist, create a new one
		newPlayer, createErr := u.q.CreatePlayer(ctx, player) // Rename the 'err' variable to 'createErr'
		if createErr != nil {
			return db.Player{}, createErr
		}
		return newPlayer, nil
	}

	// Player exists, check if the rating needs to be updated
	if existingPlayer.Rating != player.Rating {
		existingPlayer.Rating = player.Rating
		updatedPlayer, updateErr := u.q.UpdatePlayer(ctx, db.UpdatePlayerParams{
			PlayerID: existingPlayer.PlayerID,
			Rating:   player.Rating,
			Name:     existingPlayer.Name,
			ImageUrl: existingPlayer.ImageUrl,
		})
		if updateErr != nil {
			return db.Player{}, updateErr
		}
		return updatedPlayer, nil
	}

	return existingPlayer, nil
}
