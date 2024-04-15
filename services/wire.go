//go:build wireinject
// +build wireinject

package services

import (
	"github.com/corentings/chessbet/app/handlers"
	"github.com/corentings/chessbet/infrastructures"
	"github.com/corentings/chessbet/services/bet"
	"github.com/corentings/chessbet/services/match"
	"github.com/corentings/chessbet/services/tournament"
	"github.com/corentings/chessbet/services/user"
	"github.com/google/wire"
)

func InitializeUserHandler() handlers.UserController {
	wire.Build(handlers.NewUserController, user.NewUseCase, infrastructures.GetPgxConn)
	return handlers.UserController{}
}

func InitializeJwtMiddleware() handlers.JwtMiddleware {
	wire.Build(handlers.NewJwtMiddleware, user.NewUseCase, infrastructures.GetPgxConn)
	return handlers.JwtMiddleware{}
}

func InitializeBetHandler() handlers.BetController {
	wire.Build(handlers.NewBetController, bet.NewUseCase, infrastructures.GetPgxConn)
	return handlers.BetController{}
}

func InitializeTournamentHandler() handlers.TournamentController {
	wire.Build(handlers.NewTournamentController, tournament.NewUseCase, infrastructures.GetPgxConn)
	return handlers.TournamentController{}
}

func InitializeMatchHandler() handlers.MatchController {
	wire.Build(handlers.NewMatchController, match.NewUseCase, infrastructures.GetPgxConn)
	return handlers.MatchController{}
}
