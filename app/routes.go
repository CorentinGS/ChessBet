package app

import (
	"net/http"

	"github.com/corentings/chessbet/app/handlers"
	"github.com/corentings/chessbet/assets"
	"github.com/corentings/chessbet/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (i *InstanceSingleton) registerStaticRoutes(e *echo.Echo) {
	g := e.Group("/static", StaticAssetsCacheControlMiddleware)

	g.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       ".",
		Browse:     false,
		Filesystem: assets.Assets(),
	}))
}

func (i *InstanceSingleton) registerRoutes(e *echo.Echo) {
	serviceContainer := services.DefaultServiceContainer()
	user := serviceContainer.UserHandler()
	pageController := handlers.NewPageController()

	// Page routes
	e.GET("/", pageController.GetIndex)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Image routes
	usersRoutes := e.Group("/users")

	// User routes
	usersRoutes.GET("/discord/login", user.DiscordLogin)
	usersRoutes.GET("/discord/callback", user.DiscordCallback)
}
