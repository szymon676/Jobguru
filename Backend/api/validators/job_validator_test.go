package validators

import (
	"testing"

	"github.com/szymon676/jobguru/types"
)

func TestVerifyJob(t *testing.T) {
	correctJobReq := types.JobReq{
		UserID:      1,
		Title:       "Cobol developer for bank",
		Company:     "bank sa",
		Skills:      []string{"cobol", "fortran"},
		Salary:      30000,
		Description: "cobol developer to develop perfoment systems",
		Currency:    "USD",
		Date:        "2022-12-12",
		Location:    "manchaster",
	}

	err := VerifyJobReq(correctJobReq)
	if err != nil {
		t.Fatalf("verification shouldn't return error")
	}

	wrongJobReq := types.JobReq{
		UserID:      1,
		Title:       "",
		Company:     "Bank",
		Skills:      []string{"cobol", "fortran"},
		Salary:      10,
		Description: "developer",
		Currency:    "USD",
		Date:        "2022-12-12",
		Location:    "123",
	}

	err = VerifyJobReq(wrongJobReq)
	if err == nil {
		t.Fatalf("verifier should return error")
	}
}
