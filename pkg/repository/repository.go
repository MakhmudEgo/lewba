package repository

import (
	"context"
	"matcha/pkg/model"
)

type UserDB interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdatePassword(ctx context.Context, username string, password []byte) error
	UserEmail(ctx context.Context, email string) (*model.User, error)
	UserUsername(ctx context.Context, username string) (*model.User, error)
}
