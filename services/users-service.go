package services

import (
	"github.com/solidaeon/cv-api/domains/users"
	"github.com/solidaeon/cv-api/utils/errors"
)

func GetUser() {}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func FindUser() {}
