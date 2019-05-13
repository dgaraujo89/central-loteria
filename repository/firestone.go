package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

const connectionOpenError = "The connection wasn't opened correctly"

// Firestone is a struct for control connection with firestone
type Firestone struct {
	client  *firestore.Client
	context context.Context
}

// Connect open a connection with firestone
func (f *Firestone) Connect(projectID string) error {
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

// Close is a function to close a opened connection with firestone
func (f *Firestone) Close() {
	if f.client != nil {
		f.client.Close()
		f.client = nil
	}
}

// Client return a pointer for client connection
func (f *Firestone) Client() (*firestore.Client, error) {
	if f.client == nil {
		return nil, fmt.Errorf(connectionOpenError)
	}

	return f.client, nil
}

// SaveOrUpdate is a function to save or update a register
func (f *Firestone) SaveOrUpdate(collectionName string, id string, data interface{}) error {
	if f.client == nil {
		return fmt.Errorf(connectionOpenError)
	}

	_, err := f.client.Collection(collectionName).Doc(id).Set(f.context, data)

	if err != nil {
		return err
	}

	return nil
}

// Save is a function to save
func (f *Firestone) Save(collectionName string, data interface{}) error {
	if f.client == nil {
		return fmt.Errorf(connectionOpenError)
	}

	_, err := f.client.Collection(collectionName).NewDoc().Set(f.context, data)

	if err != nil {
		return err
	}

	return nil
}

// Get a list of documents in the collection
func (f *Firestone) GetAll(collectionName string) (*firestore.DocumentIterator, error) {
	if f.client == nil {
		return nil, fmt.Errorf(connectionOpenError)
	}

	return f.client.Collection(collectionName).Documents(f.context), nil
}

// Collection get a collection reference
func (f *Firestone) Collection(collectionName string) (*firestore.CollectionRef, error) {
	if f.client == nil {
		return nil, fmt.Errorf(connectionOpenError)
	}

	return f.client.Collection(collectionName), nil
}

// Delete a register from id
func (f *Firestone) Delete(collectionName string, id string) error {
	if f.client == nil {
		return fmt.Errorf(connectionOpenError)
	}

	_, err := f.client.Collection(collectionName).Doc(id).Delete(f.context)

	return err
}
