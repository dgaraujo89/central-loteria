package services

import (
	"errors"
	"log"
	"strings"

	"github.com/diegogomesaraujo/central-loteria/pkg/commons"
	"github.com/diegogomesaraujo/central-loteria/pkg/entities"
	"google.golang.org/api/iterator"
)

type UserService struct{}

const userCollectionName = "users"

func (u *UserService) FindAll() ([]entities.User, error) {
	firestore, err := firestoreConnect()
	if err != nil {
		return nil, err
	}
	defer firestore.Close()

	docIter, err := firestore.GetAll(userCollectionName)

	if err != nil {
		log.Printf("Error when get all users: %v\n", err)
		return nil, err
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

	return users, nil
}

func (u *UserService) FindByID(id string) (*entities.User, error) {
	firestore, err := firestoreConnect()
	if err != nil {
		return nil, err
	}
	defer firestore.Close()

	doc, errDoc := firestore.Get(userCollectionName, id)

	if errDoc != nil {
		if strings.Contains(errDoc.Error(), "NotFound") {
			return nil, errors.New("User not found")
		}

		return nil, errDoc
	}

	if doc == nil {
		return nil, errors.New("User not found")
	}

	var user entities.User
	doc.DataTo(&user)

	return &user, nil
}

func (u *UserService) Save(user entities.User) error {
	firestore, err := firestoreConnect()
	if err != nil {
		return err
	}
	defer firestore.Close()

	return firestore.Save(userCollectionName, user)
}

func (u *UserService) Update(user entities.User) error {
	firestore, err := firestoreConnect()
	if err != nil {
		return err
	}
	defer firestore.Close()

	return firestore.Update(userCollectionName, user.ID, user)
}

func (u *UserService) UpdatePassword(user entities.User) error {
	firestore, err := firestoreConnect()
	if err != nil {
		return err
	}
	defer firestore.Close()

	return firestore.UpdatePartial(userCollectionName, user.ID, map[string]interface{}{
		"Password": commons.Encrypt(user.Password),
	})
}

func (u *UserService) Delete(id string) error {
	firestore, err := firestoreConnect()
	if err != nil {
		return err
	}
	defer firestore.Close()

	return firestore.Delete(userCollectionName, id)
}

func (u *UserService) FindByEmail(email string) *entities.User {
	firestore, err := firestoreConnect()
	if err != nil {
		return nil
	}
	defer firestore.Close()

	collection, err := firestore.Collection(userCollectionName)

	if err != nil {
		log.Printf("Erro: %v", err)
		return nil
	}

	var user entities.User

	iter := collection.Where("email", "==", email).Documents(firestore.GetContext())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		doc.DataTo(&user)
		break
	}

	return &user
}
