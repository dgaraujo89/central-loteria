package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/diegogomesaraujo/central-loteria/pkg/commons"
	"github.com/diegogomesaraujo/central-loteria/pkg/entities"
	"github.com/diegogomesaraujo/central-loteria/pkg/services"
)

// UserController to access users resources
type UserController struct{}

const userCollectionName = "users"

// Find all users
func (c *UserController) Find(w http.ResponseWriter, r *http.Request) {
	userService := services.UserService{}

	users, err := userService.FindAll()
	if err != nil {
		log.Printf("Error when get all users: %v\n", err)
		json.NewEncoder(w).Encode([]entities.User{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// FindByID select a single user
func (c *UserController) FindByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		handleExceptionError(w, "User not found", http.StatusNotFound)
		return
	}

	userService := services.UserService{}

	user, err := userService.FindByID(params["id"])

	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			handleExceptionError(w, "User not found", http.StatusNotFound)
			return
		}

		log.Printf("An error was occured when get collection ref: %v\n", err)
		handleExceptionError(w, "An erro occured when get the user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Add a new user
func (c *UserController) Add(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if readBodyFromJSON(w, r, &user) != nil {
		return
	}

	userValidation := UserValidation{}

	if userValidation.AddValidation(w, user) != nil {
		return
	}

	userService := services.UserService{}

	user.ID = ""
	user.Password = commons.Encrypt(user.Password)

	err := userService.Save(user)
	if err != nil {
		handleExceptionError(w, "User not created", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

// Update the user
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if readBodyFromJSON(w, r, &user) != nil {
		return
	}

	if user.ID == "" {
		handleExceptionError(w, "User not found", http.StatusNotFound)
		return
	}

	userService := services.UserService{}

	err := userService.Update(user)

	if err != nil {
		log.Printf("Error when update user: %v\n", err)
		handleExceptionError(w, "User not updated", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdatePassword password to user
func (c *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if readBodyFromJSON(w, r, &user) != nil {
		return
	}

	if user.ID == "" {
		handleExceptionError(w, "User not found", http.StatusNotFound)
		return
	}

	if user.Password == "" || len(user.Password) < 6 {
		handleExceptionError(w, "Password invalid", http.StatusBadRequest)
		return
	}

	userService := services.UserService{}

	err := userService.UpdatePassword(user)

	if err != nil {
		log.Printf("An error was occured when update: %v\n", err)
		handleExceptionError(w, "Error to update the password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete remove the user from ID
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		handleExceptionError(w, "User not found", http.StatusNotFound)
		return
	}

	userService := services.UserService{}

	err := userService.Delete(params["id"])

	if err != nil {
		log.Printf("An error was occured when remove the user: %v\n", err)
		handleExceptionError(w, "Error to remove the user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
