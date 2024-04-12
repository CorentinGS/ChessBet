package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/corentings/chessbet/pkg/oauth"
	"github.com/corentings/chessbet/pkg/random"
	"github.com/corentings/chessbet/services/user"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	useCase user.IUseCase
}

const (
	oauthStateLength = 16
	// SessionTokenCookieKey is the key for the session token cookie.
	SessionTokenCookieKey = "session_token"
	ExpiresDuration       = 24 * time.Hour
)

var stateCache sync.Map //nolint:gochecknoglobals //State cache

func NewUserController(user user.IUseCase) UserController {
	return UserController{useCase: user}
}

func (uc *UserController) DiscordLogin(c echo.Context) error {
	state, _ := random.GetRandomGeneratorInstance().GenerateSecretCode(oauthStateLength)

	// Create the dynamic redirect URL for login
	redirectURL := oauth.GetDiscordURL() + "&state=" + state

	// Add the state to the cache

	stateCache.Store(state, true)

	slog.Info("Redirecting to Discord login", slog.String("redirectURL", redirectURL))
	return Redirect(c, redirectURL, http.StatusSeeOther)
}

func (uc *UserController) DiscordCallback(c echo.Context) error {
	// Get the state from the query
	state := c.QueryParam("state")

	// Check if the state is in the cache
	if _, ok := stateCache.Load(state); !ok {
		slog.Warn("State not found in cache", slog.String("state", state))
		return c.NoContent(http.StatusUnauthorized)
	}

	// Delete the state from the cache
	stateCache.Delete(state)

	// Get the code from the query
	code := c.QueryParam("code")

	// Get the access token from Discord
	accessToken, err := oauth.GetDiscordAccessToken(c.Request().Context(), code)
	if err != nil {
		slog.Error("Failed to get access token from Discord", slog.String("error", err.Error()), slog.String("code", code))
		return c.NoContent(http.StatusInternalServerError)
	}

	if accessToken == "" {
		slog.Error("Access token is empty")
		return c.NoContent(http.StatusInternalServerError)
	}

	// Get the user info from Discord
	userInfo, err := oauth.GetDiscordData(c.Request().Context(), accessToken)
	if err != nil {
		slog.Error("Failed to get user info from Discord", slog.String("error", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	var discordUser oauth.DiscordLogin

	// Unmarshal the user info into a struct
	err = json.Unmarshal([]byte(userInfo), &discordUser)
	if err != nil {
		slog.Error("Failed to unmarshal user info", slog.String("error", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	slog.Debug("Discord user info", slog.String("userInfo", userInfo))

	// Login the user
	token, err := uc.useCase.LoginOauthDiscord(c.Request().Context(), discordUser)
	if err != nil {
		slog.Error("Failed to login user", slog.String("error", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	cookie := &http.Cookie{
		Name:     SessionTokenCookieKey,
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(ExpiresDuration),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie)

	// Redirect to the home page
	return c.Redirect(http.StatusSeeOther, "/app")
}
