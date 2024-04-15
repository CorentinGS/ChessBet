package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/corentings/chessbet/app/views/page"
	"github.com/corentings/chessbet/services/tournament"
	"github.com/labstack/echo/v4"
)

type TournamentController struct {
	useCase tournament.IUseCase
}

func NewTournamentController(useCase tournament.IUseCase) TournamentController {
	return TournamentController{useCase: useCase}
}

func (tc *TournamentController) CreateTournamentFromLichessID(c echo.Context) error {
	lichessID := c.FormValue("lichess_id")
	_, err := tc.useCase.CreateTournamentFromLichessID(c.Request().Context(), lichessID)
	if err != nil {
		slog.Error("Error creating tournament from lichess id", slog.String("error", err.Error()))
		return Render(c, http.StatusForbidden, page.AdminError("Failed to create tournament from lichess ID"))
	}
	return Render(c, http.StatusForbidden, page.AdminError("Tournament created successfully"))
}

func (tc *TournamentController) GetTournamentsInProgress(c echo.Context) error {
	tournaments, err := tc.useCase.GetTournamentsInProgress(c.Request().Context())
	if err != nil {
		slog.Error("Error getting tournaments in progress", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	slog.Debug("Tournaments in progress", slog.Int("count", len(tournaments)), slog.Int("id", int(tournaments[0].TournamentID)))

	slog.Info("Tournaments in progress", slog.Int("count", len(tournaments)))

	return Render(c, http.StatusOK, page.Gallery(tournaments))
}

func (tc *TournamentController) GetTournamentByID(c echo.Context) error {
	tournamentID := c.Param("id")

	// convert string to int32
	tournamentIDInt, err := strconv.ParseInt(tournamentID, 10, 32)
	if err != nil {
		slog.Error("Error parsing tournament ID", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusBadRequest)
	}

	tournament, err := tc.useCase.GetTournamentByID(c.Request().Context(), int32(tournamentIDInt))
	if err != nil {
		slog.Error("Error getting tournament by id", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	slog.Info("Tournament by id", slog.Int64("id", tournamentIDInt), slog.String("name", tournament.Name))

	return Render(c, http.StatusOK, page.TournamentComponent(tournament))
}
