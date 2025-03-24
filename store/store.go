package store

import (
	"context"
	"database/sql"
)

type Store struct {
	User interface {
		Create(context.Context, *sql.Tx, *User) (*User, error)
		CreateAndInvite(context.Context, *User) (*User, error)
		GetStr() string
	}
}

var MockStore Store

func NewStore(db *sql.DB) Store {
	return Store{
		User: &UserStore{db},
	}
}
