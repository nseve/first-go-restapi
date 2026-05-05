package service

import (
	"errors"
	"time"

	"github.com/nseve/first-go-restapi/internal/models"
	"github.com/nseve/first-go-restapi/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

type AuthService struct {
	users     repository.UserRepository
	jwtSecret string
}

func NewAuthService(users repository.UserRepository, secret string) *AuthService {
	return &AuthService{users: users, jwtSecret: secret}
}

func (s *AuthService) Register(email, password string) error {
	_, err := s.users.GetByEmail(email)
	if err == nil {
		return ErrUserAlreadyExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{Email: email, PasswordHash: string(hash)}

	return s.users.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.users.GetByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return s.generateToken(user.ID)
}

func (s *AuthService) generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.jwtSecret))
}
