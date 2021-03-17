package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(user UserInput) (User, error)
	Login(email string, password string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(user UserInput) (User, error) {
	var newUser User

	newUser.Email = user.Email
	newUser.Name = user.Name
	newUser.Role = "user"
	newUser.Username = user.Username

	// encrypt password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return newUser, err
	}

	newUser.Password = string(passwordHash)

	// kirim lewat repo -> db
	results, err := s.repository.Save(newUser)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (s *service) Login(email string, password string) (User, error) {
	user := User{}

	// cari email
	searchEmail, err := s.repository.FindEmail(email)
	if err != nil {
		return searchEmail, err
	}

	// apabila email tidak ada
	if searchEmail.ID == 0 {
		return searchEmail, errors.New("Email Not Found!")
	}

	// apabila ada, cocokkan passwordnya
	status := bcrypt.CompareHashAndPassword([]byte(searchEmail.Password), []byte(password))

	if status != nil {
		fmt.Println("password salah")
		return user, errors.New("Wrong Password!")
	}

	// password cocok
	return searchEmail, nil

}
