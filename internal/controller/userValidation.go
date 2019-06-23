package controller

import (
	"errors"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/diegogomesaraujo/central-loteria/pkg/entities"
)

type UserValidation struct{}

func (v *UserValidation) AddValidation(w http.ResponseWriter, user entities.User) error {
	// TODO: change the e-mail validation method
	if checkmail.ValidateFormat(user.Email) != nil {
		handleExceptionError(w, "Invalid e-mail address", http.StatusBadRequest)
		return errors.New("Invalid e-mail address")
	}

	if user.Name == "" {
		handleExceptionError(w, "Invalid name", http.StatusBadRequest)
		return errors.New("Invalid name")
	}

	if user.Password == "" || len(user.Password) < 6 {
		handleExceptionError(w, "Invalid password", http.StatusBadRequest)
		return errors.New("Invalid password")
	}

	return nil
}
