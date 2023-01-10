package users

import (
	"database/sql"
	"fmt"

	"github.com/solidaeon/cv-api/datasources/mysql/users_db"
	"github.com/solidaeon/cv-api/utils/date_utils"
	"github.com/solidaeon/cv-api/utils/errors"
	"github.com/solidaeon/cv-api/utils/mysql_utils"
)

const (
	insertStmt = "insert into users_db values(?, ?, ?, ?, ?)"
	getStmt    = "select * from users_db where id = ?"
	updateStmt = "update users_db set first_name=?, last_name=?, email=? where id=?"
	deleteStmt = "delete from users_db where id=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(getStmt)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if err == sql.ErrNoRows {
			return errors.NewNotFoundError(fmt.Sprintf("user with id %d not found.", user.Id))
		}
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(insertStmt)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.Id, user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error fetching userid. %s", err.Error()))
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(updateStmt)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(deleteStmt)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}
