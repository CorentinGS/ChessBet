package services

import (
	"sync"

	"github.com/corentings/chessbet/app/handlers"
)

type ServiceContainer struct {
	userHandler handlers.UserController
	jwtHandler  handlers.JwtMiddleware
}

var (
	container *ServiceContainer //nolint:gochecknoglobals // Singleton
	once      sync.Once         //nolint:gochecknoglobals // Singleton
)

func DefaultServiceContainer() *ServiceContainer {
	imageHandler := InitializeUserHandler()
	jwtHandler := InitializeJwtMiddleware()

	return NewServiceContainer(imageHandler, jwtHandler)
}

func NewServiceContainer(userHandler handlers.UserController, jwtHandler handlers.JwtMiddleware,
) *ServiceContainer {
	once.Do(func() {
		container = &ServiceContainer{
			userHandler: userHandler,
			jwtHandler:  jwtHandler,
		}
	})
	return container
}

func (sc *ServiceContainer) UserHandler() handlers.UserController {
	return sc.userHandler
}

func (sc *ServiceContainer) JwtMiddleware() handlers.JwtMiddleware {
	return sc.jwtHandler
}

func (sc *ServiceContainer) TournamentHandler() handlers.TournamentController {
	return InitializeTournamentHandler()
}

func (sc *ServiceContainer) BetHandler() handlers.BetController {
	return InitializeBetHandler()
}

func (sc *ServiceContainer) MatchHandler() handlers.MatchController {
	return InitializeMatchHandler()
}
