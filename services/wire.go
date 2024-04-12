//go:build wireinject
// +build wireinject

package services

import (
	"github.com/corentings/chessbet/app/handlers"
	"github.com/corentings/chessbet/infrastructures"
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
