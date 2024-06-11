package authsrv

import (
	"errors"
	"time"
	"urfu-radio-journal/internal/config"
	"urfu-radio-journal/internal/models"

	"github.com/golang-jwt/jwt"
)

var (
	errBadPass    = errors.New("invalid password")
	errBadToken   = errors.New("bad token")
	errNoPayload  = errors.New("no payload")
	errBadPayload = errors.New("wrong value type in payload")
)

type AuthService struct {
	admin         models.Admin
	tokenLifetime time.Duration
	secret        []byte
}

func NewAuthService(conf config.AuthConfig) *AuthService {
	return &AuthService{
		admin: models.Admin{
			Password: conf.Password,
			Username: conf.Login,
		},
		tokenLifetime: time.Duration(conf.TokenLifetime),
		secret:        []byte(conf.Secret),
	}
}

func (as *AuthService) checkAdmin(admin models.Admin) bool {
	return admin.Username == as.admin.Username && admin.Password == as.admin.Password
}

func (as *AuthService) Login(admin models.Admin) (string, error) {
	if as.checkAdmin(admin) {
		return as.createToken(admin)
	}
	return "", errBadPass
}

func (as *AuthService) createToken(admin models.Admin) (string, error) {
	iat := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": admin.Username,
		"iat":      iat.Unix(),
		"exp":      iat.Add(time.Hour * as.tokenLifetime).Unix(), // Время жизни токена
	})
	tokenString, err := token.SignedString(as.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *AuthService) ValidateToken(tokenIn string) error {
	hashSecretGetter := func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, errBadToken
		}
		return as.secret, nil
	}
	token, err := jwt.Parse(tokenIn, hashSecretGetter)
	if err != nil || !token.Valid {
		return err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errNoPayload
	}

	username, ok := payload["username"].(string)
	if !ok {
		return errBadPayload
	}

	if as.admin.Username != username {
		return errBadToken
	}
	return nil
}
