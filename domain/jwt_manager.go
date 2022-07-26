package domain

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtManager struct {
	secretKey string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Id uint `json:"id"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JwtManager {
	return &JwtManager{
		secretKey: secretKey,
		tokenDuration: tokenDuration,
	}
}


func (m *JwtManager) Generate(u *User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().UTC().Unix(),
			ExpiresAt: time.Now().UTC().Add(m.tokenDuration).Unix(),
		},
		Id: u.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JwtManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("jwt is not okey")
			}
			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("jwt claim is not okey")
	}
	return claims, nil
}