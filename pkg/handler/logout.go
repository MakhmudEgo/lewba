package handler

import (
	"matcha/pkg/repository"
	"net/http"
)

type logout struct {
	db repository.UserDB
}

func Logout(db repository.UserDB) *logout {
	return &logout{db: db}
}

func (l *logout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}
