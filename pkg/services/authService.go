package services

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
	"urfu-radio-journal/internal/models"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	admin         models.Admin
	tokenLifetime time.Duration
	secret        []byte
}

func NewAuthService() *AuthService {
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		log.Fatal("Missing admin password in environvent variables.")
	}
	username := os.Getenv("ADMIN_USERNAME")
	if username == "" {
		log.Fatal("Missing admin username in environvent variables.")
	}
	tokenLifetime, err := strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	if err != nil {
		log.Fatal("Can't parse token lifetime.")
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Missing secret in environvent variables.")
	}
	return &AuthService{
		admin: models.Admin{
			Password: password,
			Username: username,
		},
		tokenLifetime: time.Duration(tokenLifetime),
		secret:        []byte(secret),
	}
}

func (this *AuthService) checkAdmin(admin models.Admin) bool {
	return admin.Username == this.admin.Username && admin.Password == this.admin.Password
}

func (this *AuthService) Login(admin models.Admin) (token string, err error) {
	if this.checkAdmin(admin) {
		token, err = this.CreateToken(admin)
		return
	}
	err = errors.New("Check login or password.")
	return
}

func (this *AuthService) CreateToken(admin models.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": admin.Username,
		"exp":      time.Now().Add(time.Hour * this.tokenLifetime).Unix(), // Время жизни токена
	})
	tokenString, err := token.SignedString(this.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (this *AuthService) ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return this.secret, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, exists := claims["username"].(string); exists {
			return nil
		}
	}
	return errors.New("Invalid token")
}
