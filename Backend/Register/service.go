package Register

import (
	"context"

	"github.com/Komal-0110/User-Authentication-Service/models"
)

type Service struct {
	repo UserRepo
}

func NewService(repo UserRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddUser(ctx context.Context, user models.UserReq) error {
	if err := s.repo.AddUser(ctx, &models.User{}); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUsers(ctx context.Context) ([]models.UserRes, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	userRes := make([]models.UserRes, 0)
	for _, user := range users {
		userRes = append(userRes, models.UserRes{
			Username:      user.Username,
			Email:         user.Email,
			Role:          user.Role,
			AccountStatus: user.AccountStatus,
		})
	}

	return userRes, nil
}

func (s *Service) GetUserById(ctx context.Context, userId int) (models.UserRes, error) {
	user, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return models.UserRes{}, err
	}

	userRes := models.UserRes{
		Username:      user.Username,
		Email:         user.Email,
		Role:          user.Role,
		AccountStatus: user.AccountStatus,
	}

	return userRes, nil
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
