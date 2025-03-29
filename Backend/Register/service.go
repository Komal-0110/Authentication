package Register

import (
	"context"

	"github.com/Komal-0110/User-Authentication-Service/models"
)

type Service struct {
	repo UserRepo
}

type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewService(repo UserRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddUser(ctx context.Context, user UserReq) error {
	if err := s.repo.AddUser(ctx, &models.User{}); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUserById(ctx context.Context, userId int) (models.User, error) {
	user, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user models.User) error {
	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, userId int) error {
	if err := s.repo.DeleteUser(ctx, userId); err != nil {
		return err
	}

	return nil
}
