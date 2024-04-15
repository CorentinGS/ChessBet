package handlers

import (
	"github.com/corentings/chessbet/services/match"
)

type MatchController struct {
	useCase match.IUseCase
}

func NewMatchController(useCase match.IUseCase) MatchController {
	return MatchController{useCase: useCase}
}
