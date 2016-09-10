package membersimport_test

import (
	"github.com/supercoopbdx/membersimport"
	"testing"
)

var mailTestCases = []struct {
	firstName string
	lastName  string
	expected  string
}{
	{"Randy", "Marsh", "randy.marsh" + membersimport.SUPERCOOP_MAIL_SUFFIX},
	{"ELIOT", "MAC MAHON-MOH", "eliot.macmahonmoh" + membersimport.SUPERCOOP_MAIL_SUFFIX},
	{"Jean-Edouard", "schmidt", "jean-edouard.schmidt" + membersimport.SUPERCOOP_MAIL_SUFFIX},
	{"  Jean Edouard  ", "schmidt", "jean-edouard.schmidt" + membersimport.SUPERCOOP_MAIL_SUFFIX},
}

func TestSupercoopMail(t *testing.T) {
	// Arrange
	target := membersimport.SupercoopMail{}

	for index, testCase := range mailTestCases {
		// Act
		actual := target.Generate(testCase.firstName, testCase.lastName)
		// Assert
		if actual != testCase.expected {
			t.Errorf("SupercoopMail(%d): expected %s, actual %s", index, testCase.expected, actual)
		}
	}
}
