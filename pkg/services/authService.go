package services

import (
	"errors"
	"os"
	"strconv"
	"urfu-radio-journal/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

type AuthService struct {
	store cookie.Store
}

func NewAuthService() *AuthService {
	secret := os.Getenv("SECRET")
	cookieMaxAge, _ := strconv.Atoi(os.Getenv("COOKIE_MAX_AGE"))
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		MaxAge:   cookieMaxAge, // seconds
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // only for HTTPS
	})

	return &AuthService{store: store}
}

func checkAdmin(admin models.Admin) bool {
	password := os.Getenv("ADMIN_PASSWORD")
	username := os.Getenv("ADMIN_USERNAME")
	return admin.Username == username && admin.Password == password
}

func (this *AuthService) Login(admin models.Admin, session sessions.Session) error {
	if checkAdmin(admin) {
		session.Set("admin", admin.Username)
		if err := session.Save(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unathorized")
}

func (this *AuthService) Logout(session sessions.Session) error {
	if admin := session.Get("admin"); admin != nil {
		session.Delete("admin")
		if err := session.Save(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unathorized")
}

func (this *AuthService) GetStore() cookie.Store {
	return this.store
}
