package service

import (
	"ecommerce/repository"
)

type User interface {
	GetAll()
}

func NewUser(repo repository.User) User {
	return &user{
		repo: repo,
	}
}

type user struct {
	repo repository.User
}

func (srv *user) GetAll() {

}