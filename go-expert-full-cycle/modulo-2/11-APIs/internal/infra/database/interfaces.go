package database

import "github.com/mllcarvalho/golang/go-expert-full-cycle/modulo-2/11-APIs/internal/entity"

type UserInterface interface {
	Create(user entity.User) error
	findByEmail(email string) (*entity.User, error)
}
