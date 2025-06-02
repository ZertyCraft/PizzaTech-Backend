package services

import (
	"errors"
	"time"

	"pizzatech/config"
	"pizzatech/internal/domain/models"
	"pizzatech/internal/domain/repositories"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(email, pass string, role models.Role) error
	Login(email, pass string) (string, error)
}

type authService struct {
	users     repositories.UserRepository
	jwtSecret string
}

func NewAuthService(u repositories.UserRepository, cfg *config.Config) AuthService {
	return &authService{users: u, jwtSecret: cfg.JWTSecret}
}

func (s *authService) Register(email, pass string, role models.Role) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user := &models.User{Email: email, Password: string(hash), Role: role}
	return s.users.Create(user)
}

func (s *authService) Login(email, pass string) (string, error) {
	user, err := s.users.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return "", errors.New("invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}
