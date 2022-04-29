package repository

import (
	"context"
	"simple-projects/models"
)

func (q *SQLStore) GetTransactions(ctx context.Context, email string) ([]*models.Transaction, error) {

	m, err := models.Transactions().All(ctx, q.db)
	if err != nil {
		return nil, nil
	}

	return m, err

}
