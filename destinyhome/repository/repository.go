package repository

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

type Repository interface {
	GetUser(string) (*user, error)
}

type repository struct {
	client  *datastore.Client
	timeout time.Duration
}

func NewRepository(projectID string, timeout time.Duration) (Repository, error) {

	const operation = "NewRepository"

	// Create context for datastore.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create datastore client.
	c, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, errors.Wrap(err, operation)
	}

	return &repository{c, timeout}, nil
}

func (r *repository) GetUser(uname string) (*user, error) {

	const operation = "repository.GetUser"

	// Create context for datastore.
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	// Create a query to find a user with the same username.
	query := datastore.NewQuery("User").Filter("Username =", uname).Limit(1)

	// Run the query.
	var users = []user{}
	if _, err := r.client.GetAll(ctx, query, &users); err != nil {
		return nil, errors.Wrap(err, operation)
	}

	// Check if the user was found.
	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return &users[0], nil
}
