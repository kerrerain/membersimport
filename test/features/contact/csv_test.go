package contact_test

import (
	target "github.com/supercoopbdx/membersimport/features/contact"
	"testing"
)

var testCases = []struct {
	Input          *target.Contact
	ExpectedOutput target.Contact
}{
	{
		&target.Contact{
			FirstName: "John",
			LastName:  "Doe",
			Phone:     " 06 01 02 03   04",
			HelloAsso: " X  ",
			Address:   "1, rue henri marceille 33140 VILLENAVE D'ORNON",
		},
		target.Contact{
			FirstName:       "John",
			LastName:        "Doe",
			Phone:           "0601020304",
			HelloAsso:       "true",
			Address:         "1, rue henri marceille 33140 VILLENAVE D'ORNON",
			Address_ZipCode: "33140",
			Address_City:    "VILLENAVE D'ORNON",
			Address_Street:  "1, rue henri marceille",
		},
	},
	{
		&target.Contact{
			Address: " 33000 BORDEAUX",
		},
		target.Contact{
			HelloAsso:       "false",
			Address:         " 33000 BORDEAUX",
			Address_ZipCode: "33000",
			Address_City:    "BORDEAUX",
		},
	},
	{
		&target.Contact{
			Address: "3 aux près de Mède",
		},
		target.Contact{
			HelloAsso:      "false",
			Address:        "3 aux près de Mède",
			Address_Street: "3 aux près de Mède",
		},
	},
}

func TestProcessRecord(t *testing.T) {
	// Arrange
	for index, testCase := range testCases {
		// Act
		target.ProcessRecord(testCase.Input)
		actual := *testCase.Input
		// Assert
		if actual != testCase.ExpectedOutput {
			t.Errorf("TestCase(%d): expected %s, actual %s", index, testCase.ExpectedOutput, actual)
		}
	}
}
