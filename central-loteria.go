package main

import (
	"log"
	"os"

	"github.com/diegogomesaraujo/central-loteria/entities"

	"google.golang.org/api/iterator"

	"github.com/diegogomesaraujo/central-loteria/repository"
)

func main() {
	firestone := repository.Firestone{}

	err := firestone.Connect(os.Getenv("GCLOUD_PROJECT_ID"))
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
		return
	}

	defer firestone.Close()

	saveUser(&firestone)
	listUsers(&firestone)
	//deleteUser(&firestone)
}

func listUsers(firestone *repository.Firestone) {
	iter, _ := firestone.GetAll("users")

	var user entities.User

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		doc.DataTo(&user)

		user.ID = doc.Ref.ID

		log.Printf("User: %v", user)
	}
}

func saveUser(firestone *repository.Firestone) {
	user := entities.User{
		Name:     "User 1",
		Email:    "user1@gmail.com",
		Password: "123456",
	}

	firestone.Save("users", user)
}

func deleteUser(firestone *repository.Firestone) {
	firestone.Delete("users", "kcjzawZk8RZcDrSxboWj")
}
