package user

import (
	"github.com/simiken/golang-cleanarchitecture/entity"
)

func NewInteractor(repository Repository) *Interactor {
	return &Interactor{
		Repository: repository,
	}
}

type Interactor struct {
	Repository Repository
}

func (i *Interactor) CreateUser(email string, password string) error {
	u := entity.NewUser(email, password)
	return i.Repository.Create(&u)
}

func (i *Interactor) FindUserByID(id int) (*entity.User, error) {
	u, err := i.Repository.FindByID(id)
	return u, err
}
