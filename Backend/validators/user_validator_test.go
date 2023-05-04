package validators

import (
	"testing"

	"github.com/szymon676/jobguru/types"
)

func TestVerifyRegister(t *testing.T) {
	wrongRegister := types.RegisterUser{
		Name:     "An",
		Email:    "a@gmal.com",
		Password: "421",
	}

	if err := VerifyRegisterReq(wrongRegister); err == nil {
		t.Errorf("expected error have no error")
	}

	correctRegister := types.RegisterUser{
		Name:     "Ana",
		Email:    "ana@gmail.com",
		Password: "4321",
	}

	if err := VerifyRegisterReq(correctRegister); err != nil {
		t.Errorf("expected no error got %c", err)
	}
}
