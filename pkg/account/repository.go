package account

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetAccountById(ID string) (*Account, error)
	CreateAccount(Name string) (*Account, error)
}

type repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return repository{db}
}

func (r repository) GetAccountById(id string) (*Account, error) {
	acc := &Account{}

	_, err := uuid.FromString(id)

	if err != nil {
		return nil, fmt.Errorf("%s is not uuid", id)
	}

	err = r.db.QueryRow(context.Background(), "SELECT id, name FROM accounts WHERE id = $1", id).Scan(&acc.ID, &acc.Name)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("no account with id=%s found", id)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get account with id=%s: %w", id, err)
	}

	return acc, nil
}

func (r repository) CreateAccount(name string) (*Account, error) {
	acc := &Account{}

	err := r.db.QueryRow(context.Background(), "INSERT INTO accounts(name) VALUES ($1) RETURNING id, name", name).Scan(&acc.ID, &acc.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to create account with name=%s: %w", name, err)
	}

	return acc, nil
}
