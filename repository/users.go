package repository

import (
	"context"
	"fmt"
	"simple-projects/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (q *SQLStore) GetUser(ctx context.Context, email string) (*models.User, error) {

	m, err := models.Users(
		qm.Where("email = ?", email),
	).One(ctx, q.db)
	if err != nil {
		return nil, nil
	}

	return m, err

}

func (q *SQLStore) GetUsers(ctx context.Context) ([]*models.User, error) {

	m, err := models.Users().All(ctx, q.db)
	if err != nil {
		return nil, nil
	}

	return m, err

}

func (q *SQLStore) AuthUser(ctx context.Context, email string) ([]*models.User, error) {

	m, err := models.Users().All(ctx, q.db)
	if err != nil {
		return nil, nil
	}

	return m, err

}

func (q *SQLStore) CreateUser(ctx context.Context, mUser *models.User) error {
	if err := mUser.Insert(ctx, q.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to create mUser: %v", err)
	}
	return nil

}
