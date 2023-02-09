package auth

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/seggga/hranilka-auth/internal/domain/models"
	"github.com/seggga/hranilka-auth/internal/ports"
)

// Service implements main auth logic
type Service struct {
	db  ports.UserStorage
	jwt jwtCfg
}

type jwtCfg struct {
	secret   string
	duration int
}

// New creates a new auth service
func New(db ports.UserStorage, secret string, duration int) *Service {
	return &Service{
		db: db,
		jwt: jwtCfg{
			secret:   secret,
			duration: duration,
		},
	}
}

// Validate checks token provided
func (s *Service) Validate(ctx context.Context, token string) (string, error) {
	select {
	case <-ctx.Done():
		return "", errors.New("context has been closed")
	default:
	}

	res, login := checkToken(token, s.jwt.secret)
	if !res {
		return "", errors.New("bad token")
	}
	return login, nil
}

// Login checks login/password correctness and
// produces token
func (s *Service) Login(ctx context.Context, login, password string) (*models.Token, error) {
	user, err := s.db.Get(ctx, login)
	if err != nil {
		return nil, err
	}

	// check password correctness
	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password))
	if err != nil {
		return nil, err
	}

	// generate token
	token, err := createToken(login, s.jwt.secret, s.jwt.duration)
	if err != nil {
		return nil, fmt.Errorf("cannot generate JWT: %v", err)
	}

	return &models.Token{Access: token}, nil
}

// SignUp registers a new user
func (s *Service) SignUp(ctx context.Context, user *models.User) error {
	if user.Login == "" {
		return ErrEmptyLogin
	}
	if len(user.Password) > 72 { // limitation of bcrypt algorythm
		return ErrPassTooLong
	}
	cost := 10
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return fmt.Errorf("cannot generate hash from given pass: %w", err)
	}

	return s.db.Create(ctx, &models.User{
		Name:     user.Name,
		Login:    user.Login,
		Password: "",
		PassHash: string(passHash),
		Email:    user.Email,
	})
}
