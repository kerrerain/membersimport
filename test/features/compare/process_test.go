package compare_test

import (
	"github.com/supercoopbdx/membersimport/features/adhesion"
	target "github.com/supercoopbdx/membersimport/features/compare"
	"github.com/supercoopbdx/membersimport/features/contact"
	"reflect"
	"testing"
)

var testCases = []struct {
	Name           string
	Contacts       []*contact.Contact
	Accessions     []*adhesion.Member
	ExpectedOutput []string
}{
	{
		"Should not display an error if the data is correct.",
		[]*contact.Contact{
			&contact.Contact{
				Id:        125,
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		[]*adhesion.Member{
			&adhesion.Member{
				Id:        125,
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		[]string{},
	},
	{
		"Should display an error if the data is missing.",
		[]*contact.Contact{
			&contact.Contact{
				Id:        125,
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		[]*adhesion.Member{
			&adhesion.Member{
				Id:        128,
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		[]string{"Le contact pour l'adhésion de John Doe 128 n'existe pas."},
	},
	{
		"Should display an error if the data does not match.",
		[]*contact.Contact{
			&contact.Contact{
				Id:        125,
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		[]*adhesion.Member{
			&adhesion.Member{
				Id:        125,
				FirstName: "Herman",
				LastName:  "Schwarz",
			},
		},
		[]string{"Différence de nom pour ID 125 : Herman Schwarz dans adhésions, John Doe dans contacts."},
	},
}

func TestProcessRecord(t *testing.T) {
	// Arrange
	for _, testCase := range testCases {
		// Act
		actual := target.CompareContactsAndAccessions(testCase.Contacts, testCase.Accessions)
		// Assert
		if !reflect.DeepEqual(actual, testCase.ExpectedOutput) {
			t.Errorf("TestCase: (%s)\n expected %+v\n actual %+v\n",
				testCase.Name, testCase.ExpectedOutput, actual)
		}
	}
}
