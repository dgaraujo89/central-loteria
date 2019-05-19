package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/badoux/checkmail"

	"google.golang.org/api/iterator"

	"github.com/diegogomesaraujo/central-loteria/entities"
	"github.com/diegogomesaraujo/central-loteria/utils"
)

// UserController to access users resources
type UserController struct{}

const userCollectionName = "users"

// Find all users
func (c *UserController) Find(w http.ResponseWriter, r *http.Request) {
	firestore, err := firestoreConnect(w)
	if err != nil {
		log.Printf("Error when connect to firestore: %v\n", err)
		return
	}

	defer firestore.Close()

	docIter, err := firestore.GetAll(userCollectionName)

	if err != nil {
		log.Printf("Error when get all users: %v\n", err)
		json.NewEncoder(w).Encode([]entities.User{})
		return
	}

	var user entities.User
	var users []entities.User

	for {
		doc, err := docIter.Next()
		if err == iterator.Done {
			break
		}

		doc.DataTo(&user)
		user.ID = doc.Ref.ID
		users = append(users, user)
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

	firestore, err := firestoreConnect(w)
	if err != nil {
		log.Printf("Error when connect to firestore: %v\n", err)
		return
	}

	defer firestore.Close()

	doc, errDoc := firestore.Get(userCollectionName, params["id"])

	if errDoc != nil {
		if strings.Contains(errDoc.Error(), "NotFound") {
			handleExceptionError(w, "User not found", http.StatusNotFound)
			return
		}

		log.Printf("An error was occured when get collection ref: %v\n", errDoc)
		handleExceptionError(w, "An erro occured when get the user", http.StatusInternalServerError)
		return
	}

	if doc == nil {
		handleExceptionError(w, "User not found", http.StatusNotFound)
		return
	}

	var user entities.User
	doc.DataTo(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Add a new user
func (c *UserController) Add(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if readBodyFromJSON(w, r, &user) != nil {
		return
	}

	// TODO: change the e-mail validation method
	if checkmail.ValidateFormat(user.Email) != nil {
		handleExceptionError(w, "Invalid e-mail address", http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		handleExceptionError(w, "Invalid name", http.StatusBadRequest)
		return
	}

	if user.Password == "" || len(user.Password) < 6 {
		handleExceptionError(w, "Invalid password", http.StatusBadRequest)
		return
	}

	user.ID = ""
	user.Password = commons.Encrypt(user.Password)

	firestore, err := firestoreConnect(w)
	if err != nil {
		log.Printf("Error when connect to firestore: %v\n", err)
		return
	}
	defer firestore.Close()

	err = firestore.Save(userCollectionName, user)
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

	firestore, err := firestoreConnect(w)
	if err != nil {
		log.Printf("Error when connect to firestore: %v\n", err)
		return
	}
	defer firestore.Close()

	err = firestore.Update(userCollectionName, user.ID, user)
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

	firestore, err := firestoreConnect(w)
	if err != nil {
		log.Printf("Error when connect to firestore: %v\n", err)
		return
	}

	defer firestore.Close()

	err = firestore.UpdatePartial(userCollectionName, user.ID, map[string]interface{}{
		"Password": commons.Encrypt(user.Password),
	})

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

	firestore, err := firestoreConnect(w)
	if err != nil {
		log.Printf("Error when connect to firestore: %v\n", err)
		return
	}

	defer firestore.Close()

	err = firestore.Delete(userCollectionName, params["id"])

	if err != nil {
		log.Printf("An error was occured when remove the user: %v\n", err)
		handleExceptionError(w, "Error to remove the user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
