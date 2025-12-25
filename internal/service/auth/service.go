package auth

import (
	"github.com/yourname/music-shop/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterRequest) error
	Login(input LoginRequest) (string, error)
}

type AuthService struct {
	users user.Repository
}

func NewAuthService(users user.Repository) *AuthService {
	return &AuthService{users: users}
}

func (s *AuthService) Register(input RegisterRequest) error {
	// check user exists
	_, err := s.users.FindByEmail(input.Email)
	if err == nil {
		return ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := &user.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hash),
		Role:     user.RoleUser,
	}

	return s.users.Save(u)
}

func (s *AuthService) Login(input LoginRequest) (string, error) {
	u, err := s.users.FindByEmail(input.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)) != nil {
		return "", ErrInvalidCredentials
	}

	return GenerateToken(u.ID, string(u.Role))
}
