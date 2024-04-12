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

	index := page.IndexPage("ChessBet", "", false, GetNonce(c), hero)

	return Render(c, http.StatusOK, index)
}

func (p *PageController) GetHome(c echo.Context) error {
	hero := page.Home()

	home := page.HomePage("ChessBet", "Home", true, GetNonce(c), hero)

	return Render(c, http.StatusOK, home)
}
