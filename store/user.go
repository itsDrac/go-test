package store

import (
	"context"
	"database/sql"
	"time"
)

type UserStore struct {
	db *sql.DB
}

type password struct {
	hashed []byte
	text   string
}

type User struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   password  `json:"-"`
	isVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) (*User, error) {
	// TODO: Check if user exist with same email.
	// TODO: If user exist with same email, return an error.
	// TODO: If user doesn't exist Insert user in database.
	// TODO: Update user pointer with CreatedAt field.
	// TODO: return pointer to user.
	return nil, nil
}

func (s *UserStore) CreateAndInvite(ctx context.Context, user *User) (*User, error) {
	// TODO: Create new sql.Tx and name it tx.
	// TODO: Call Create function using tx and user.
	// TODO: If no error is returned by Create function, then call SendInvite function.
	// TODO: if no error is returned by SendInvite function, then commit the changes.
	// TODO: return pointer to user.
	return nil, nil
}
