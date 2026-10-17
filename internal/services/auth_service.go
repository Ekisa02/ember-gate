package services

import (
	"errors"
	"time"

	"github.com/example/ember-gate/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB        *gorm.DB
	JWTSecret string
}

func NewAuthService(db *gorm.DB, secret string) *AuthService {
	return &AuthService{DB: db, JWTSecret: secret}
}

func (s *AuthService) Register(email, password string) (*models.User, error) {
	var exists models.User
	if err := s.DB.Where("email = ?", email).First(&exists).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{Email: email, PasswordHash: string(hash)}
	if err := s.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Authenticate(email, password string) (string, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	signed, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}
	return signed, nil
}
