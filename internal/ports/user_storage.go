package ports

import (
	"context"

	"github.com/seggga/hranilka-auth/internal/domain/models"
)

type UserStorage interface {
	Get(ctx context.Context, login string) (*models.User, error)
}
