package membersimport

import (
	"github.com/supercoopbdx/membersimport"
	"testing"
)

/*
	MOCKS
*/

type SupercoopMailMock struct {
	GenerateCalled bool
}

func (s *SupercoopMailMock) Generate(firstName string, lastName string) string {
	s.GenerateCalled = true

	if len(firstName) == 0 || len(lastName) == 0 {
		return ""
	} else {
		return "dummy.response@supercoop.fr"
	}
}

/*
	TESTS
*/

func TestProcessRecord(t *testing.T) {
	// Arrange
	generator := &SupercoopMailMock{}

	record := membersimport.Member{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "06 01 02 03   04",
	}

	// Act
	result := membersimport.ProcessRecord(generator, record)

	// Assert
	if result.SupercoopMail != "dummy.response@supercoop.fr" {
		t.Error("ProcessRecord: expected SupercoopMail to have been processed")
	}
	if result.Phone != "0601020304" {
		t.Error("ProcessRecord: expected Phone to have been processed")
	}
}
