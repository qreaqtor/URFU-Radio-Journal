package services

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"urfu-radio-journal/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

type AuthService struct {
	store cookie.Store
}

func NewAuthService(domain string) *AuthService {
	secret := os.Getenv("SECRET")
	secureStr := os.Getenv("SECURE_HTTPS")
	secure, err := strconv.ParseBool(secureStr)
	if err != nil {
		log.Fatalf("Can't parse secure parametr: %s", err.Error())
	}
	httpOnlyStr := os.Getenv("HTTP_ONLY")
	httpOnly, err := strconv.ParseBool(httpOnlyStr)
	if err != nil {
		log.Fatalf("Can't parse httpOnly parametr: %s", err.Error())
	}
	cookieMaxAge, _ := strconv.Atoi(os.Getenv("COOKIE_MAX_AGE"))
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		MaxAge:   cookieMaxAge, // seconds
		Path:     "/",
		HttpOnly: httpOnly,
		Secure:   secure, // only for HTTPS
		SameSite: http.SameSiteNoneMode,
		Domain:   domain,
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
	return errors.New("Check login or password.")
}

func (this *AuthService) Logout(session sessions.Session) error {
	if admin := session.Get("admin"); admin != nil {
		session.Delete("admin")
		if err := session.Save(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Unathorized. Missing cookie.")
}

func (this *AuthService) GetStore() cookie.Store {
	return this.store
}
