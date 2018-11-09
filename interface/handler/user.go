package handler

import (
	"fmt"
	"net/http"

	usergw "github.com/simiken/golang-cleanarchitecture/interface/gateway/user"
	useruc "github.com/simiken/golang-cleanarchitecture/usecase/user"
)

func GetUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := 1
		repo := usergw.NewRepository()
		usecase := useruc.NewInteractor(repo)
		u, err := usecase.FindUserByID(id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}
}

func PostUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		email := "email"
		password := "password"
		repo := usergw.NewRepository()
		usecase := useruc.NewInteractor(repo)
		err := usecase.CreateUser(email, password)
		if err != nil {
			fmt.Println(err)
		}
	}
}
