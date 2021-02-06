package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

//CheckUserType renews the user tokens when they login
func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return err
}
