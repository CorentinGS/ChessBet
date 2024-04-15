package handlers

import (
	"github.com/corentings/chessbet/services/bet"
)

type BetController struct {
	useCase bet.IUseCase
}

func NewBetController(useCase bet.IUseCase) BetController {
	return BetController{useCase: useCase}
}
