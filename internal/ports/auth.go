package ports

import (
	"context"

	"github.com/seggga/hranilka-auth/internal/domain/models"
)

type Auther interface {
	Validate(ctx context.Context, token string) (string, error)
	Login(ctx context.Context, login, password string) (*models.Token, error)
	SignUp(ctx context.Context, user *models.User) error
	ChangeProfile(ctx context.Context, oldLogin string, user *models.User) error
	ChangePass(ctx context.Context, login, newPass string) error
}
