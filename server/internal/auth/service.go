package auth

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req RegisterRequest) (AuthResponse, error)
	Login(req LoginRequest) (AuthResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(req RegisterRequest) (AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, err
	}

	user := User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := s.repo.Create(user)

	if err != nil {
		return AuthResponse{}, err
	}

	token, err := GenerateToken(createdUser.ID, createdUser.Email)
	if err != nil {
		return AuthResponse{}, err
	}

	return builtAuthResponse(createdUser, token), nil

}

func (s *service) Login(req LoginRequest) (AuthResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return AuthResponse{}, ErrInvalidCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return AuthResponse{}, ErrInvalidCredential
	}

	token, err := GenerateToken(user.ID, user.Email)
	if err != nil {
		return AuthResponse{}, err
	}

	return builtAuthResponse(user, token), nil
}

func builtAuthResponse(user User, token string) AuthResponse {
	response := AuthResponse{
		Token: token,
	}
	response.User.ID = user.ID
	response.User.Name = user.Name
	response.User.Email = user.Email
	return response
}
