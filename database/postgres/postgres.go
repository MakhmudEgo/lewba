package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"matcha/pkg/model"
	"matcha/pkg/repository"
	"os"
)

type postgres struct {
	db *pgx.Conn
}

func Postgres(ctx context.Context) repository.UserDB {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	println("db connect")
	return &postgres{db: conn}
}

func (db *postgres) CreateUser(ctx context.Context, user *model.User) error {
	_, err := db.db.Exec(ctx, "insert into users(username, password) values($1, $2)", user.Username, user.Password)
	return err
}
func (db *postgres) UpdatePassword(ctx context.Context, username string, password []byte) error {
	return nil

}
func (db *postgres) UserEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil

}
func (db *postgres) UserUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

//todo::??
func (db *postgres) Close() {
	err := db.db.Close(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}
