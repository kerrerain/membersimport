package contact_test

import (
	target "github.com/supercoopbdx/membersimport/features/contact"
	"testing"
)

var testCases = []struct {
	Name           string
	Input          *target.Contact
	ExpectedOutput target.Contact
}{
	{
		"Should split the address into ZipCode, City and Street.",
		&target.Contact{
			FirstName: "John",
			LastName:  "Doe",
			Phone:     " 06 01 02 03   04",
			HelloAsso: " X  ",
			Address:   "1, rue john doe 33000 VAGRANT D'ULIEL",
		},
		target.Contact{
			FirstName:       "John",
			LastName:        "Doe",
			Phone:           "0601020304",
			HelloAsso:       "true",
			Address:         "1, rue john doe 33000 VAGRANT D'ULIEL",
			Address_ZipCode: "33000",
			Address_City:    "VAGRANT D'ULIEL",
			Address_Street:  "1, rue john doe",
			Type:            "CONTACT",
		},
	},
	{
		"Should split the address even if there are missing parts (Street).",
		&target.Contact{
			Address: " 33000 BORDEAUX",
		},
		target.Contact{
			HelloAsso:       "false",
			Address:         " 33000 BORDEAUX",
			Address_ZipCode: "33000",
			Address_City:    "BORDEAUX",
			Type:            "CONTACT",
		},
	},
	{
		"Should split the address even if there are missing parts (ZipCode, City).",
		&target.Contact{
			Address: "3 rue des nuages",
		},
		target.Contact{
			HelloAsso:      "false",
			Address:        "3 rue des nuages",
			Address_Street: "3 rue des nuages",
			Type:           "CONTACT",
		},
	},
	{
		"Should extract the date of contact from the event.",
		&target.Contact{
			Event: "14/04/2016 - Réunion d'info",
		},
		target.Contact{
			HelloAsso:     "false",
			Event:         "14/04/2016 - Réunion d'info",
			DateOfContact: "14/04/2016",
			Type:          "CONTACT",
		},
	},
	{
		"Should do nothing if the date of contact is already set.",
		&target.Contact{
			Event:         "Réunion d'info",
			DateOfContact: "14/04/2016",
		},
		target.Contact{
			HelloAsso:     "false",
			Event:         "Réunion d'info",
			DateOfContact: "14/04/2016",
			Type:          "CONTACT",
		},
	},
	{
		"Should set the type to ADHERENT if there is a member id.",
		&target.Contact{
			Id:      583,
			Address: "33000 BORDEAUX",
		},
		target.Contact{
			Id:              583,
			HelloAsso:       "false",
			Address:         "33000 BORDEAUX",
			Address_ZipCode: "33000",
			Address_City:    "BORDEAUX",
			Type:            "ADHERENT",
		},
	},
}

func TestProcessRecord(t *testing.T) {
	// Arrange
	for _, testCase := range testCases {
		// Act
		target.ProcessRecord(testCase.Input)
		actual := *testCase.Input
		// Assert
		if actual != testCase.ExpectedOutput {
			t.Errorf("TestCase: (%s)\n expected %+v\n actual %+v\n",
				testCase.Name, testCase.ExpectedOutput, actual)
		}
	}
}
