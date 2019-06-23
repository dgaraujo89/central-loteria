package controller

import (
	"net/http"
	"io"

	"github.com/diegogomesaraujo/central-loteria/pkg/commons"
	"github.com/diegogomesaraujo/central-loteria/pkg/entities"
	"github.com/diegogomesaraujo/central-loteria/pkg/services"
)

type LoginController struct{}

func (l *LoginController) Auth(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if readBodyFromJSON(w, r, &user) != nil {
		return
	}

	if user.Email == "" {
		handleExceptionError(w, "User not found", http.StatusNotFound)
		return
	}

	userServices := services.UserService{}

	userDb := userServices.FindByEmail(user.Email)

	password := commons.Encrypt(user.Password)

	if userDb.Password != password {
		handleExceptionError(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	io.WriteString(w, `{"token": "8n734v87v2638bfn9d432n7yf3d873d8n75476b"}`)
}
