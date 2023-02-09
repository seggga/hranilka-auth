package ports

import (
	"context"

	"github.com/seggga/hranilka-auth/internal/domain/models"
)

type Auther interface {
	Validate(ctx context.Context, token string) (string, error)
	Login(ctx context.Context, login, password string) (*models.Token, error)
}
