package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Komal-0110/User-Authentication-Service/Register"
	"github.com/uptrace/bun"
)

var ErrNotFound = errors.New("no user found")

type DB struct {
	db *bun.DB
}

func NewDB(db *bun.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d *DB) AddUser(ctx context.Context, user Register.User) error {
	if err := d.db.
		NewInsert().
		Model(&user).
		Scan(ctx); err != nil {
		return err
	}

	return nil
}

func (d *DB) GetUsers(ctx context.Context) ([]Register.User, error) {
	var users []Register.User
	if err := d.db.
		NewSelect().
		Model(&users).
		Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return users, nil
}

func (d *DB) GetUser(ctx context.Context, userId int) (Register.User, error) {
	var user Register.User
	if err := d.db.
		NewSelect().
		Model(&user).
		Where("Id = ?", userId).
		Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return Register.User{}, ErrNotFound
		}
		return Register.User{}, err
	}

	return user, nil
}

func (d *DB) UpdateUser(ctx context.Context, user Register.User) error {
	if err := d.db.
		NewUpdate().
		Model(&user).
		Where("Id = ?", user.Id).
		Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (d *DB) DeleteUser(ctx context.Context, userId int) error {
	var user Register.User
	if err := d.db.NewDelete().Model(&user).Where("id = ?", userId).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	return nil
}
