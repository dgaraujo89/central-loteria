package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/api/iterator"

	"github.com/diegogomesaraujo/central-loteria/entities"
	"github.com/diegogomesaraujo/central-loteria/exception"
	"github.com/diegogomesaraujo/central-loteria/repository"
)

type UserController struct{}

// Find all users
func (c *UserController) Find(w http.ResponseWriter, r *http.Request) {
	firestone := repository.Firestone{}
	err := firestone.Connect()

	if err != nil {
		exception.HandleError(w, exception.Exception{"Connection error", http.StatusInternalServerError})
		return
	}

	defer firestone.Close()

	docIter, err := firestone.GetAll("users")

	if err != nil {
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

	json.NewEncoder(w).Encode(users)
}

// Add a new user
func (c *UserController) Add(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if r.Header.Get("Content-Type") != "application/json" {
		exception.HandleError(w, exception.Exception{"Content-Type not supported", http.StatusBadRequest})
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Erro when try read body")
		exception.HandleError(w, exception.Exception{"Erro when try read body", http.StatusInternalServerError})
		return
	}

	if !json.Valid(body) {
		exception.HandleError(w, exception.Exception{"Bad Request", http.StatusBadRequest})
		return
	}

	json.Unmarshal(body, &user)

	log.Printf("New user: %v\n", user)

}
