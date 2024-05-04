package authsrv

import (
	"errors"
	"time"
	"urfu-radio-journal/internal/models"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	admin         models.Admin
	tokenLifetime time.Duration
	secret        []byte
}

func NewAuthService(password, username, secret string, tokenLifetime int) *AuthService {
	return &AuthService{
		admin: models.Admin{
			Password: password,
			Username: username,
		},
		tokenLifetime: time.Duration(tokenLifetime),
		secret:        []byte(secret),
	}
}

func (as *AuthService) checkAdmin(admin models.Admin) bool {
	return admin.Username == as.admin.Username && admin.Password == as.admin.Password
}

func (as *AuthService) Login(admin models.Admin) (token string, err error) {
	if as.checkAdmin(admin) {
		token, err = as.createToken(admin)
		return
	}
	err = errors.New("check login or password")
	return
}

func (as *AuthService) createToken(admin models.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": admin.Username,
		"exp":      time.Now().Add(time.Hour * as.tokenLifetime).Unix(), // Время жизни токена
	})
	tokenString, err := token.SignedString(as.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *AuthService) ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return as.secret, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, exists := claims["username"].(string); exists {
			return nil
		}
	}
	return errors.New("invalid token")
}
