package user

import (
	"database/sql"

	"github.com/simiken/golang-cleanarchitecture/entity"
	"github.com/simiken/golang-cleanarchitecture/infrastructure/db"
	"github.com/simiken/golang-cleanarchitecture/usecase/user"
)

func NewRepository() user.Repository {
	return Repository{db: db.DB}
}

type Repository struct {
	db *sql.DB
}

func (r Repository) Create(user *entity.User) error {
	return nil
}

func (r Repository) FindAll() ([]*entity.User, error) {
	return nil, nil
}

func (r Repository) FindByID(id int) (*entity.User, error) {
	return nil, nil
}
