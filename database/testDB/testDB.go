package testDB

import (
	"context"
	"matcha/pkg/model"
	"matcha/pkg/repository"
)

type testDb struct {
}

func TestDb(ctx context.Context) repository.UserDB {
	return &testDb{}
}

func (db testDb) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}
func (db testDb) UpdatePassword(ctx context.Context, username string, password []byte) error {
	return nil
}
func (db testDb) UserEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}
func (db testDb) UserUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}
