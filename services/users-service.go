package services

import (
	"github.com/solidaeon/cv-api/domains/users"
	"github.com/solidaeon/cv-api/utils/errors"
)

func GetUser() {}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUser() {}
