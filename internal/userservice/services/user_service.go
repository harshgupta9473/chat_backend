package services

import (
	"context"
	"errors"
	"github.com/harshgupta9473/chatapp/internal/userservice/dto"
	"github.com/harshgupta9473/chatapp/internal/userservice/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Signup(ctx context.Context, req *dto.SignupRequest) error {
	// Check if user exists
	_, err := s.userRepo.GetUserByMobile(ctx, req.Mobile)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &dto.User{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Password: string(hashedPassword),
	}

	return s.userRepo.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.User, error) {
	user, err := s.userRepo.GetUserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, errors.New("invalid mobile or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid mobile or password")
	}

	return user, nil
}
