package services

import (
	"errors"
	"urfu-radio-journal/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

type AuthService struct {
	store cookie.Store
}

func NewAuthService() *AuthService {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   3600, // seconds
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // only for HTTPS
	})

	return &AuthService{store: store}
}

func checkAdmin(admin models.Admin) bool {
	return admin.Username == "admin" && admin.Password == "admin"
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
