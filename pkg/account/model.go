package account

import "github.com/gofrs/uuid/v5"

type Account struct {
	ID   uuid.UUID
	Name string
}
