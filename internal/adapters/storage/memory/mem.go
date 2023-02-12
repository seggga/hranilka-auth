package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	er "github.com/seggga/hranilka-auth/internal/adapters/storage"
	"github.com/seggga/hranilka-auth/internal/domain/models"
)

type Storage struct {
	m     map[string]models.User
	mutex sync.Mutex
}

// New creates a new memory storage
func New() *Storage {
	m := make(map[string]models.User, 2)
	// m["user1"] = models.User{
	// 	ID:       uuid.New(),
	// 	Name:     "",
	// 	Login:    "test-user1",
	// 	PassHash: "",
	// 	Email:    "test@test.com",
	// }

	return &Storage{
		m:     m,
		mutex: sync.Mutex{},
	}
}

// Get looks for a user by login
func (s *Storage) Get(ctx context.Context, login string) (*models.User, error) {
	select {
	case <-ctx.Done():
		return nil, er.ClosedContext
	default:
		s.mutex.Lock()
		u, ok := s.m[login]
		s.mutex.Unlock()
		if !ok {
			return nil, er.NotFound
		}
		return &u, nil
	}
}

// Create adds a new user to the storage
func (s *Storage) Create(ctx context.Context, user *models.User) error {
	select {
	case <-ctx.Done():
		return er.ClosedContext
	default:
		u := models.User{
			ID:       uuid.New(),
			Name:     user.Name,
			Login:    user.Login,
			PassHash: user.PassHash,
			Email:    user.Email,
		}

		s.mutex.Lock()
		s.m[u.Login] = u
		s.mutex.Unlock()

		return nil
	}
}

// Set sets user's profile data exept PassHash
func (s *Storage) Set(ctx context.Context, user *models.User) error {
	select {
	case <-ctx.Done():
		return er.ClosedContext
	default:
		u := models.User{
			ID:    uuid.New(),
			Name:  user.Name,
			Login: user.Login,
			Email: user.Email,
		}

		s.mutex.Lock()
		s.m[u.Login] = u
		s.mutex.Unlock()

		return nil
	}
}

// SetPass sets a new password
func (s *Storage) SetPass(ctx context.Context, login, passHash string) error {
	select {
	case <-ctx.Done():
		return er.ClosedContext
	default:
		s.mutex.Lock()
		u := s.m[login]
		u.PassHash = passHash
		s.m[login] = u
		s.mutex.Unlock()

		return nil
	}
}
