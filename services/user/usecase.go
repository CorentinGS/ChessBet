package user

import (
	"context"
	"errors"

	db "github.com/corentings/chessbet/db/sqlc"
	"github.com/corentings/chessbet/pkg/jwt"
	"github.com/corentings/chessbet/pkg/oauth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UseCase struct {
	q *db.Queries
}

func NewUseCase(dbConn *pgxpool.Pool) IUseCase {
	q := db.New(dbConn)

	return &UseCase{q: q}
}

func (u *UseCase) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return u.q.GetUser(ctx, id)
}

func (u *UseCase) GetUsers(ctx context.Context) ([]db.User, error) {
	return u.q.GetUsers(ctx)
}

func (u *UseCase) CreateUser(ctx context.Context, user db.CreateUserParams) (db.User, error) {
	return u.q.CreateUser(ctx, user)
}

func (u *UseCase) LoginOauthDiscord(ctx context.Context, oauth oauth.DiscordLogin) (string, error) {
	// Get the user from the database by the email
	user, err := u.q.GetUserByEmail(ctx, oauth.Email)
	if err != nil {
		// If the user does not exist, register the user
		if errors.Is(err, pgx.ErrNoRows) {
			user, err = u.RegisterOauth(ctx, db.CreateUserParams{
				EmailAddress: oauth.Email,
				Username:     oauth.Username,
				Points:       0,
			})
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	if user.EmailAddress != oauth.Email {
		return "", errors.New("email address does not match")
	}

	// If the oauth user is different from the user in the database, update the user
	if user.Username != oauth.Username {
		user, err = u.q.UpdateUser(ctx, db.UpdateUserParams{
			UserID:   user.UserID,
			Username: oauth.Username,
		})
		if err != nil {
			return "", err
		}
	}

	token, err := jwt.GetJwtInstance().GetJwt().GenerateToken(ctx, user.UserID)
	if err != nil {
		return "", err
	}

	// Return the user
	return token, nil
}

func (u *UseCase) RegisterOauth(ctx context.Context, oauth db.CreateUserParams) (db.User, error) {
	// Check if the user already exists
	user, err := u.q.GetUserByEmail(ctx, oauth.EmailAddress)
	if err != nil {
		// If the user does not exist, create a new user
		if errors.Is(err, pgx.ErrNoRows) {
			user, err = u.q.CreateUser(ctx, oauth)
			if err != nil {
				return db.User{}, err
			}
		} else {
			return db.User{}, err
		}
	}

	return user, nil
}
