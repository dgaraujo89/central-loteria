package services

import (
	"errors"
	"log"

	"github.com/diegogomesaraujo/central-loteria/pkg/repository"
)

func firestoreConnect() (repository.Firestore, error) {
	firestore := repository.Firestore{}
	err := firestore.Connect()

	if err != nil {
		log.Printf("Error when connect to firestore: %v\n", err)
		return firestore, errors.New("Connection error")
	}

	return firestore, nil
}
