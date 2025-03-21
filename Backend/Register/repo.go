package Register

import (
	"context"

	"github.com/Komal-0110/User-Authentication-Service/models"
	"github.com/Komal-0110/User-Authentication-Service/sqlite"
)

type UserRepo interface {
	AddUser(ctx context.Context, user *models.User) error
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, userId int) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, userId int) error
}

type Users struct {
	db sqlite.DB
}

func NewUserRepo(db sqlite.DB) *Users {
	return &Users{
		db: db,
	}
}

func (r *Users) RegisterUser(ctx context.Context, user models.User) error {
	return r.db.AddUser(ctx, user)
}

func (r *Users) GetUsers(ctx context.Context) ([]models.User, error) {
	return r.db.GetUsers(ctx)
}

func (r *Users) GetUserById(ctx context.Context, userId int) (models.User, error) {
	return r.db.GetUser(ctx, userId)
}

func (r *Users) UpdateUser(ctx context.Context, user models.User) error {
	return r.db.UpdateUser(ctx, user)
}

func (r *Users) DeleteUser(ctx context.Context, userId int) error {
	return r.DeleteUser(ctx, userId)
}
