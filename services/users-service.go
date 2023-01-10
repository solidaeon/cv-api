package services

import (
	"fmt"

	"github.com/solidaeon/cv-api/domains/users"
	"github.com/solidaeon/cv-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {

	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	userOnRecord, err := GetUser(user.Id)

	fmt.Println("on record")
	fmt.Println(userOnRecord)

	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			userOnRecord.FirstName = user.FirstName
		}
		if user.LastName != "" {
			userOnRecord.LastName = user.LastName
		}
		if user.Email != "" {
			userOnRecord.Email = user.Email
		}
		if user.Status != "" {
			userOnRecord.Status = user.Status
		}
	} else {
		userOnRecord.FirstName = user.FirstName
		userOnRecord.LastName = user.LastName
		userOnRecord.Email = user.Email
		userOnRecord.Status = user.Status
	}

	if err := userOnRecord.Validate(); err != nil {
		return nil, err
	}

	if err := userOnRecord.Update(); err != nil {
		return nil, err
	}

	return userOnRecord, nil
}

func DeleteUser(user users.User) *errors.RestErr {

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func FindByStatus(status string) ([]users.User, *errors.RestErr) {

	dao := &users.User{}

	return dao.FindByStatus(status)
}
