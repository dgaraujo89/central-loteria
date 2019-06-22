package repository

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
)

const connectionOpenError = "The connection wasn't opened correctly"

// Firestore is a struct for control connection with firestore
type Firestore struct {
	client  *firestore.Client
	context context.Context
}

// Connect open a connection with firestore
func (f *Firestore) Connect() error {
	projectID := os.Getenv("GCLOUD_PROJECT_ID")

	if projectID == "" {
		return fmt.Errorf("projectId not found")
	}

	f.context = context.Background()

	var err error
	f.client, err = firestore.NewClient(f.context, projectID)

	if err != nil {
		return err
	}

	return nil
}

// Close is a function to close a opened connection with firestore
func (f *Firestore) Close() {
	if f.client != nil {
		f.client.Close()
		f.client = nil
	}
}

// Client return a pointer for client connection
func (f *Firestore) Client() (*firestore.Client, error) {
	if f.client == nil {
		return nil, fmt.Errorf(connectionOpenError)
	}

	return f.client, nil
}

// GetContext return the context instance
func (f *Firestore) GetContext() context.Context {
	return f.context
}

// Update is a function to save or update a register
func (f *Firestore) Update(collectionName string, id string, data interface{}) error {
	if f.client == nil {
		return fmt.Errorf(connectionOpenError)
	}

	_, err := f.client.Collection(collectionName).Doc(id).Set(f.context, data)

	if err != nil {
		return err
	}

	return nil
}

// UpdatePartial is a function to save or update a register
func (f *Firestore) UpdatePartial(collectionName string, id string, data interface{}) error {
	if f.client == nil {
		return fmt.Errorf(connectionOpenError)
	}

	_, err := f.client.Collection(collectionName).Doc(id).Set(f.context, data, firestore.MergeAll)

	if err != nil {
		return err
	}

	return nil
}

// Save is a function to save
func (f *Firestore) Save(collectionName string, data interface{}) error {
	if f.client == nil {
		return fmt.Errorf(connectionOpenError)
	}

	_, err := f.client.Collection(collectionName).NewDoc().Set(f.context, data)

	if err != nil {
		return err
	}

	return nil
}

// GetAll a list of documents in the collection
func (f *Firestore) GetAll(collectionName string) (*firestore.DocumentIterator, error) {
	if f.client == nil {
		return nil, fmt.Errorf(connectionOpenError)
	}

	return f.client.Collection(collectionName).Documents(f.context), nil
}

// Get a single document in the collection by id
func (f *Firestore) Get(collectionName string, id string) (*firestore.DocumentSnapshot, error) {
	if f.client == nil {
		return nil, fmt.Errorf(connectionOpenError)
	}

	docSnap, err := f.client.Collection(collectionName).Doc(id).Get(f.context)
	if err != nil {
		return nil, err
	}

	return docSnap, nil
}

// Collection get a collection reference
func (f *Firestore) Collection(collectionName string) (*firestore.CollectionRef, error) {
	if f.client == nil {
		return nil, fmt.Errorf(connectionOpenError)
	}

	return f.client.Collection(collectionName), nil
}

// Delete a register from id
func (f *Firestore) Delete(collectionName string, id string) error {
	if f.client == nil {
		return fmt.Errorf(connectionOpenError)
	}

	_, err := f.client.Collection(collectionName).Doc(id).Delete(f.context)

	return err
}
