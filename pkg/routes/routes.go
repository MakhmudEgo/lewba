package routes

import (
	"matcha/pkg/handler"
	"matcha/pkg/repository"
	"net/http"
)

func init() {

	//http.Handle("refresh")
	//http.Handle("register")
}

func Init(db repository.UserDB) {
	http.Handle("/register", handler.Register(db))
	http.Handle("/login", handler.Login(db))
	http.Handle("/logout", handler.Logout(db))

	http.Handle("/register/username", handler.TestHandler())
}
