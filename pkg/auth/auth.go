package auth

type IAuth interface {
	CreateAuth() error
	DeleteAuth() error
}
