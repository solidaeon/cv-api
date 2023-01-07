package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solidaeon/cv-api/domains/users"
	"github.com/solidaeon/cv-api/services"
	"github.com/solidaeon/cv-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invald json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
