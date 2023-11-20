package account

import (
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	CreateAccount(Name string) (Account, error)
	GetAccount(ID string) (Account, error)
}

type repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return repository{db}
}

func (r repository) CreateAccount(name string) (Account, error) {
	return Account{ID: "1", Name: name}, nil
}

func (r repository) GetAccount(id string) (Account, error) {
	return Account{ID: id, Name: "hello"}, nil
}
