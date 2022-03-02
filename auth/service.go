package auth

import (
	"errors"
)

type Service interface {
	Login(input LoginInput) (DaftarPetugas, error)
	LupaPassword(input LoginInput) (DaftarPetugas, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Login(input LoginInput) (DaftarPetugas, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Email belum terdaftar")
	}

	if user.Pass != input.Password {
		return user, errors.New("Password anda salah!")
	}

	return user, nil
}

func (s *service) LupaPassword(input LoginInput) (DaftarPetugas, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Email belum terdaftar")
	}

	user.Pass = input.Password

	user, err = s.repository.LupaPassword(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
