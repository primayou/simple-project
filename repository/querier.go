package repository

import (
	"context"
	"simple-projects/models"
)

type Querier interface {
	GetUser(ctx context.Context, email string) ([]*models.User, error)
}
