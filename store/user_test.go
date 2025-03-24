package store

import (
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	// TODO: Get mockStore here.
	log.Printf("In user test function")
	t.Run("Pass Create User", func(t *testing.T) {
		// _ := mockStore.User.Create()
		got := MockStore.User.GetStr()
		want := "Working"
		if got == want {
			t.Error("Not Working")
		}
	})
	// TODO: get mock store and test the user's create function.
}
