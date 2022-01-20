package handler

import (
	"matcha/pkg/repository"
	"net/http"
)

type login struct {
}

func Login(db repository.UserDB) *login {
	return &login{}
}

func (l *login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}
