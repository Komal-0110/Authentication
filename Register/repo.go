package Register

import (
	"errors"
)

var errUserExists = errors.New("user already exist")
var errUserNotFound = errors.New("user not found")

// Add context while switching to database setup

type userRepo interface {
	RegisterUser(user *User) error
	GetUsers() ([]User, error)
	GetUserById(userId int) (User, error)
	UpdateUser(user User) error
	DeleteUser(userId int) error
}

type Users struct {
	Users []User
}

func NewUserRepo(user User) *Users {
	return &Users{
		Users: make([]User, 0),
	}
}

func (r *Users) RegisterUser(user *User) error {
	err := r.userExists(user.Username)
	if err != nil {
		return err
	}

	userLen := len(r.Users)

	if userLen == 0 {
		user.Id = 1
	} else {
		lastId := r.Users[userLen-1].Id
		user.Id = lastId + 1
	}

	r.Users = append(r.Users, *user)

	return nil
}

func (r *Users) GetUsers() ([]User, error) {
	users := make([]User, len(r.Users))
	copy(users, r.Users)

	return users, nil
}

func (r *Users) GetUserById(userId int) (User, error) {
	for _, user := range r.Users {
		if user.Id == userId {
			return user, nil
		}
	}

	return User{}, errUserNotFound
}

func (r *Users) UpdateUser(updatedUser User) error {
	for _, user := range r.Users {
		if user.Id == updatedUser.Id {
			user = updatedUser
			return nil
		}
	}

	return errUserNotFound
}

func (r *Users) DeleteUser(userId int) error {
	for idx, user := range r.Users {
		if user.Id == userId {
			r.Users = append(r.Users[:idx], r.Users[idx+1:]...)
			return nil
		}
	}

	return errUserNotFound
}

func (r *Users) userExists(userName string) error {
	for _, user := range r.Users {
		if user.Username == userName {
			return errUserExists
		}
	}

	return nil
}
