package service

import "matcha/pkg/model"

type Service interface {
	Do(user *model.User) (interface{}, error)
	Exec(data []byte) (interface{}, error)
}
