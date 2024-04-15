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

func (u *UseCase) GetUpcomingMatchByTournament(ctx context.Context, tournamentID int32) ([][]db.GetUpcomingMatchesByTournamentRow, error) {
	matches, err := u.q.GetUpcomingMatchesByTournament(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	// return the matches
	return storeMatchesByRound(matches), nil
}

// storeMatchesByRound to store matches by round name.
func storeMatchesByRound(matches []db.GetUpcomingMatchesByTournamentRow) [][]db.GetUpcomingMatchesByTournamentRow {
	// Initialize a slice to store matches by round name
	var matchesByRound [][]db.GetUpcomingMatchesByTournamentRow

	// Iterate through the matches and store them by round name
	var currentRound string
	currentRoundMatches := make([]db.GetUpcomingMatchesByTournamentRow, 0, len(matches))
	for _, match := range matches {
		if match.RoundName != currentRound {
			if len(currentRoundMatches) > 0 {
				matchesByRound = append(matchesByRound, currentRoundMatches)
			}
			currentRound = match.RoundName
			currentRoundMatches = make([]db.GetUpcomingMatchesByTournamentRow, 0, len(matches))
		}
		currentRoundMatches = append(currentRoundMatches, match)
	}

	// Add the last round's matches
	if len(currentRoundMatches) > 0 {
		matchesByRound = append(matchesByRound, currentRoundMatches)
	}

	return matchesByRound
}
