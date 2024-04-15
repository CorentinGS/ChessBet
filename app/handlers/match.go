package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/corentings/chessbet/app/views/page"
	"github.com/corentings/chessbet/services/match"
	"github.com/labstack/echo/v4"
)

type MatchController struct {
	useCase match.IUseCase
}

func NewMatchController(useCase match.IUseCase) MatchController {
	return MatchController{useCase: useCase}
}

func (mc *MatchController) GetUpcomingMatchByTournament(c echo.Context) error {
	tournamentID := c.Param("id")

	// convert string to int32
	tournamentIDInt, err := strconv.ParseInt(tournamentID, 10, 32)
	if err != nil {
		slog.Error("Error parsing tournament ID", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusBadRequest)
	}

	matches, err := mc.useCase.GetUpcomingMatchByTournament(c.Request().Context(), int32(tournamentIDInt))
	if err != nil {
		slog.Error("Error getting upcoming matches by tournament", slog.String("error", err.Error()))
		return RedirectToErrorPage(c, http.StatusInternalServerError)
	}

	slog.Info("Upcoming matches by tournament", slog.Int("count", len(matches)))

	return Render(c, http.StatusOK, page.MatchGallery(matches))
}
