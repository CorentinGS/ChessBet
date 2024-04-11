package services

import (
	"sync"

	"github.com/corentings/chessbet/app/handlers"
)

type ServiceContainer struct {
	userHandler handlers.UserController
}

var (
	container *ServiceContainer //nolint:gochecknoglobals // Singleton
	once      sync.Once         //nolint:gochecknoglobals // Singleton
)

func DefaultServiceContainer() *ServiceContainer {
	imageHandler := InitializeUserHandler()

	return NewServiceContainer(imageHandler)
}

func NewServiceContainer(userHandler handlers.UserController,
) *ServiceContainer {
	once.Do(func() {
		container = &ServiceContainer{
			userHandler: userHandler,
		}
	})
	return container
}

func (sc *ServiceContainer) UserHandler() handlers.UserController {
	return sc.userHandler
}
