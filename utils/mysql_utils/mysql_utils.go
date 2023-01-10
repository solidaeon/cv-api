package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/solidaeon/cv-api/utils/errors"
)

const (
	errNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFoundError("no user found for the given id")
		}
		return errors.NewInternalServerError("error processing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("email already exists"))
	}

	return errors.NewInternalServerError("error processing response")
}
