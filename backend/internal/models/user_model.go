package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id           int64     `json:"id"`
	FullName     string    `json:"full_name"`
	BusinessName string    `json:"business_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedOn    time.Time `json:"created_on"`
	UpdatedOn    time.Time `json:"updated_on"`
}

type CreateUser struct {
	FullName     string `json:"full_name"`
	BusinessName string `json:"business_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

func (data *CreateUser) Validate() error {
	if data.FullName == "" {
		return errors.New("full_name is required")
	}
	if data.BusinessName == "" {
		return errors.New("business_name is required")
	}
	if data.Email == "" {
		return errors.New("email is required")
	}
	if data.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type UserModel struct {
	Db *pgxpool.Pool
}

func (model *UserModel) GetById(id int64) (User, error) {
	sql := "SELECT id, full_name, business_name, email, created_on, updated_on FROM users WHERE id=$1"
	args := []any{id}

	var user User
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(
		&user.Id,
		&user.FullName,
		&user.BusinessName,
		&user.Email,
		&user.CreatedOn,
		&user.UpdatedOn,
	); err != nil {
		if err == pgx.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func (model *UserModel) Create(data CreateUser) (int64, error) {
	if err := data.Validate(); err != nil {
		return -1, err
	}

	sql := "INSERT INTO users(full_name, business_name, email, password) VALUES($1, $2, $3, $4) RETURNING id"
	args := []any{data.FullName, data.BusinessName, data.Email, data.Password}

	var insertId int64
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(&insertId); err != nil {
		return -1, err
	}

	return insertId, nil
}
