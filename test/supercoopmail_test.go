package main_test

import (
	main "github.com/supercoopbdx/membersimport"
	"testing"
)

var mailTestCases = []struct {
	firstName string
	lastName  string
	expected  string
}{
	{"Randy", "Marsh", "randy.marsh" + main.SUPERCOOP_MAIL_SUFFIX},
	{"ELIOT", "MAHON-MOH", "eliot.mahon" + main.SUPERCOOP_MAIL_SUFFIX},
	{"Hortense", "Laval Lemans", "hortense.laval" + main.SUPERCOOP_MAIL_SUFFIX},
	{"Jean-Edouard", "schmidt", "jean-edouard.schmidt" + main.SUPERCOOP_MAIL_SUFFIX},
	{"  Jean Edouard  ", "schmidt", "jean-edouard.schmidt" + main.SUPERCOOP_MAIL_SUFFIX},
}

func TestSupercoopMail(t *testing.T) {
	// Arrange
	target := main.SupercoopMail{}

	for index, testCase := range mailTestCases {
		// Act
		actual := target.Generate(testCase.firstName, testCase.lastName)
		// Assert
		if actual != testCase.expected {
			t.Errorf("SupercoopMail(%d): expected %s, actual %s", index, testCase.expected, actual)
		}
	}
}
