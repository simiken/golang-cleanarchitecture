package user

import (
	"github.com/simiken/golang-cleanarchitecture/entity"
)

type Repository interface {
	Create(*entity.User) error
	FindAll() ([]*entity.User, error)
	FindByID(id int) (*entity.User, error)
}
