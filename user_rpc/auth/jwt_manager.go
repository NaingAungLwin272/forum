package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
)

type JWTManager struct {
	secretKey     string
}

type UserClaims struct {
	jwt.StandardClaims
	Id    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{secretKey}
}

func (manager *JWTManager) Generate(user *models.DBUser, tokenDuration time.Duration) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
		},
		Id:    user.Id.Hex(),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(manager.secretKey))
}

// Verify
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, bool, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			//token ka byte type nae lar tal... ae dar ko bytetype nae check tal
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexcepted token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, false, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, false, fmt.Errorf("invalid token claims")
	}

	return claims, ok, nil
}
