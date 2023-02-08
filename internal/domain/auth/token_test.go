package auth

import (
	"errors"
	"testing"
	"time"
)

var (
	testSecret string = "good-secret"
	testLogin  string = "some-user"
)

func TestCreateToken(t *testing.T) {
	type testSet struct {
		name     string
		secret   string
		err      error
		duration int
	}

	sets := []testSet{
		{
			name:     "valid token, expect no error",
			secret:   "good-secret",
			err:      nil,
			duration: 10,
		},
		{
			name:     "token with zero duration, expect not nil error",
			secret:   "good-secret",
			err:      errors.New("zero duration"),
			duration: 0,
		},
		{
			name:     "token with empty secret word, expect not nil error",
			secret:   "",
			err:      errors.New("empty secret"),
			duration: 10,
		},
	}

	for _, v := range sets {
		_, err := createToken(testLogin, v.secret, v.duration)
		if err != nil && v.err == nil || err == nil && v.err != nil {
			t.Errorf("wrong createToken result: %s, %v", v.name, err)
		}
	}
}

func TestCheckToken(t *testing.T) {

	type testSet struct {
		name      string
		token     string
		secret    string
		validness bool
		duration  int
	}

	sets := []testSet{
		{
			name:      "valid token, expect true",
			secret:    "good-secret",
			validness: true,
			duration:  10,
		},
		{
			name:      "token out of date, expect false",
			secret:    "good-secret",
			validness: false,
			duration:  1,
		},
		{
			name:      "token with other secret word, expect false",
			secret:    "bad-secret",
			validness: false,
			duration:  10,
		},
	}
	// generate tokens
	for i, _ := range sets {
		token, _ := createToken(testLogin, sets[i].secret, sets[i].duration)
		sets[i].token = token
	}
	// wait for 1-second token expiration
	time.Sleep(time.Second)

	// check tokens
	for _, v := range sets {
		check, _ := checkToken(v.token, testSecret)
		if check != v.validness {
			t.Errorf("wrong checkToken result: %s, got %v", v.name, check)
		}
	}
}
