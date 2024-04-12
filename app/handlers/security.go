package handlers

import (
	"errors"
	"net/http"

	db "github.com/corentings/chessbet/db/sqlc"
	"github.com/corentings/chessbet/domain"
	"github.com/corentings/chessbet/pkg/jwt"
	"github.com/corentings/chessbet/services/user"
	"github.com/labstack/echo/v4"
)

// JwtMiddleware is the controller for the jwt routes.
type JwtMiddleware struct {
	user.IUseCase
}

// NewJwtController creates a new jwt controller.
func NewJwtMiddleware(user user.IUseCase) JwtMiddleware {
	return JwtMiddleware{IUseCase: user}
}

func (j *JwtMiddleware) AuthorizeUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return j.IsConnectedMiddleware(domain.PermissionUser, next)(c)
	}
}

func (j *JwtMiddleware) IsConnectedMiddleware(_ domain.Permission, next echo.HandlerFunc) func(c echo.Context) error {
	return func(c echo.Context) error {
		// get the token from the cookie
		cookie, err := c.Cookie("session_token")
		if err != nil {
			_ = RedirectToErrorPage(c, http.StatusUnauthorized)
		}

		token := cookie.Value

		userID, err := jwt.GetJwtInstance().GetJwt().GetConnectedUserID(c.Request().Context(), token)
		if err != nil {
			_ = RedirectToErrorPage(c, http.StatusUnauthorized)
		}

		userModel, err := j.IUseCase.GetUserByID(c.Request().Context(), userID)
		if err != nil {
			_ = RedirectToErrorPage(c, http.StatusUnauthorized)
		}

		// Set userModel in locals
		SetUserToContext(c, userModel)
		return next(c)
	}
}

func SetUserToContext(c echo.Context, user db.User) {
	c.Set("user", user)
}

func GetUserFromContext(c echo.Context) (db.User, error) {
	if c.Get("user") == nil {
		return db.User{}, errors.New("user is not initialized")
	}
	return c.Get("user").(db.User), nil
}
