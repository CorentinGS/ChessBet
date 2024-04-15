package handlers

import (
	"net/http"

	"github.com/corentings/chessbet/app/views/page"
	"github.com/labstack/echo/v4"
)

type PageController struct{}

func NewPageController() *PageController {
	return &PageController{}
}

func (p *PageController) GetIndex(c echo.Context) error {
	hero := page.Index()

	index := page.IndexPage("ChessBet", false, GetNonce(c), hero)

	return Render(c, http.StatusOK, index)
}

func (p *PageController) GetHome(c echo.Context) error {
	user, err := GetUserFromContext(c)
	if err != nil {
		return RedirectToErrorPage(c, http.StatusUnauthorized)
	}

	hero := page.Home(user.Username)

	home := page.HomePage("ChessBet", true, GetNonce(c), hero)

	return Render(c, http.StatusOK, home)
}

func (p *PageController) GetAdmin(c echo.Context) error {
	hero := page.Admin()

	admin := page.AdminPage("ChessBet", true, GetNonce(c), hero)

	return Render(c, http.StatusOK, admin)
}

func (p *PageController) GetTournament(c echo.Context) error {
	tournamentID := c.Param("id")

	hero := page.Tournament(tournamentID)

	tournament := page.TournamentPage("ChessBet", true, GetNonce(c), hero)

	return Render(c, http.StatusOK, tournament)
}
