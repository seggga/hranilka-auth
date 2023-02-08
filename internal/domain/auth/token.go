package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims ...
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// checkToken verifies if provided token is valid (not timedout, compromized or changed)
func checkToken(tokenString, jwtSecret string) (bool, string) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return false, ""
	}

	return token.Valid, claims.Username
}

// createToken ...
func createToken(login, secret string, duration int) (string, error) {
	// check for empty secret
	if secret == "" {
		return "", ErrEmptySecret
	}

	// check for incorrect token's time-to-live timeout
	if duration <= 0 {
		return "", ErrZeroDuration
	}

	mySigningKey := []byte(secret)

	// generate access token
	claims := &Claims{
		Username: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Second)),
			Issuer:    "hranilka-auth",
			Subject:   login,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
