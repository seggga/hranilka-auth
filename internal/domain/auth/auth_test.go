package auth

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/seggga/hranilka-auth/internal/domain/models"
)

var (
	mockStorage *MockStorage
	s           *Service
	ctx         context.Context

	testAuthLogin  = "test-login"
	testAuthSecret = "test-secret"
)

func TestMain(m *testing.M) {
	mockStorage = &MockStorage{}
	s = New(mockStorage, testAuthSecret, 1)
	ctx = context.TODO()

	os.Exit(m.Run())
}

func TestValidate(t *testing.T) {
	type testSet struct {
		name      string
		secret    string
		expectErr error
		duration  int
	}

	sets := []testSet{
		{
			name:      "valid token, expect no error",
			secret:    testAuthLogin,
			expectErr: nil,
			duration:  10,
		},
		{
			name:      "token with zero duration, expect not nil error",
			secret:    testAuthLogin,
			expectErr: errors.New("zero duration"),
			duration:  0,
		},
		{
			name:      "token with empty secret word, expect not nil error",
			secret:    "",
			expectErr: errors.New("empty secret"),
			duration:  10,
		},
	}

	for _, v := range sets {
		token, _ := createToken(testAuthLogin, v.secret, v.duration)
		user, err := s.Validate(ctx, token)
		if err != nil && v.expectErr == nil || err == nil && v.expectErr != nil {
			t.Errorf("wrong Validate result: %s, %v", v.name, err)
		}

		if err == nil && user != testAuthLogin {
			t.Errorf("got wrong login: expected %s, got %s", testAuthLogin, user)
		}
	}
}

func TestLogin(t *testing.T) {

	// Login(ctx, login, password string) (*models.Token, error)
	t.Skip()
}

type MockStorage struct{}

func (m *MockStorage) Get(ctx context.Context, login string) (*models.User, error) {
	if login == "good-login" {
		return &models.User{
			ID:    uuid.New(),
			Login: login,
		}, nil
	}

	return nil, errors.New("user with specified login was not found")

}
