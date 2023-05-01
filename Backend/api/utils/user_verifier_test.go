package utils

import (
	"testing"

	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"
)

func TestVerifyRegister(t *testing.T) {
	store, _ := storage.NewPostgreStorage("some dsn, soon...")
	userStore := storage.NewPostgreUserStorage(store)
	vs := NewVerifier(userStore)
	wrongRegister := types.RegisterUser{
		Name:     "An",
		Email:    "a@gmal.com",
		Password: "421",
	}

	if err := vs.VerifyRegister(wrongRegister); err == nil {
		t.Errorf("expected error have no error")
	}

	correctRegister := types.RegisterUser{
		Name:     "Ana",
		Email:    "ana@gmail.com",
		Password: "4321",
	}

	if err := vs.VerifyRegister(correctRegister); err != nil {
		t.Errorf("expected no error got %c", err)
	}
}
