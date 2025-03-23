package store

import (
	"testing"
)

func TestCreate(t *testing.T) {
	// TODO: Get mockStore here.
	t.Run("Pass Create User", func(t *testing.T) {
		_ := mockStore.User.Create()
	})
	// TODO: get mock store and test the user's create function.
}
