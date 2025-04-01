package querys

import (
	"errors"
	"fmt"

	"github.com/redrum15/prueba/src/db"
	"github.com/redrum15/prueba/src/db/models"
	"github.com/redrum15/prueba/src/utils"
)

func GetUserFromEmailNPassword(email, password string) (*models.User, error) {
	var user models.User

	result := db.GetDBInstance().First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	fmt.Println("passwrod de la db", user.Password)
	fmt.Println("passord de la req", password)
	var samePassword bool = utils.ComparePasswords(user.Password, []byte(password))
	if !samePassword {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}
