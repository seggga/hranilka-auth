package ports

import (
	"context"

	"github.com/seggga/hranilka-auth/internal/domain/models"
)

type UserStorage interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, login string) (*models.User, error)
	Set(ctx context.Context, user *models.User) error
}
