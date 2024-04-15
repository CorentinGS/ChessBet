package handlers

import (
	"github.com/corentings/chessbet/services/tournament"
)

type TournamentController struct {
	useCase tournament.IUseCase
}

func NewTournamentController(useCase tournament.IUseCase) TournamentController {
	return TournamentController{useCase: useCase}
}
